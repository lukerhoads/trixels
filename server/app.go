package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router  *mux.Router
	DB      *gorm.DB
	devChan chan string
}

func (a *App) Initialize(db *gorm.DB, devChan chan string) {
	a.devChan = devChan
	a.DB = db
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (r *App) initializeRoutes() {
	r.Router.HandleFunc("/trixel/{trixelId}", r.RedirectTrixel).Methods("GET")
	r.Router.HandleFunc("/pixel/{coord}", r.GetPixel).Methods("GET")
	r.Router.HandleFunc("/pixels", r.GetPixels).Methods("GET")
	r.Router.HandleFunc("/pixels", r.UpdatePixel).Methods("POST")
	r.Router.HandleFunc("/trigger/mint", r.TriggerMint).Methods("GET")
	r.Router.HandleFunc("/trigger/reset", r.ResetDB).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Println("Listening...")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))
}

func (r *App) GetPixel(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	coord := vars["coord"]
	xCoord, yCoord := DecodePixelValues(coord)
	pixel := Pixel{
		X: xCoord,
		Y: yCoord,
	}
	pixel.GetPixel(r.DB)
	json.NewEncoder(res).Encode(pixel)
}

func (r *App) GetPixels(res http.ResponseWriter, req *http.Request) {
	var pixels Pixels
	pixels.GetPixels(r.DB)
	json.NewEncoder(res).Encode(pixels)
}

type ServerError struct {
	message string
}

func (r *App) UpdatePixel(res http.ResponseWriter, req *http.Request) {
	var getPixel Pixel
	getPixel.GetPixel(r.DB)
	validUpdate := getPixel.UpdatedAt.Add(PixelUpdateTime).Before(time.Now())
	if !validUpdate {
		json.NewEncoder(res).Encode(ServerError{
			message: "Cannot update, pixel has been updated in the last 5 minutes",
		})
		return
	}

	var pixel Pixel
	err := json.NewDecoder(req.Body).Decode(&pixel)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	pixel.UpdatePixel(r.DB)
	fmt.Fprintf(res, "Success!")
}

func (r *App) TriggerMint(res http.ResponseWriter, req *http.Request) {
	r.devChan <- "mint"
	fmt.Fprintf(res, "Success!")
}

func (r *App) ResetDB(res http.ResponseWriter, req *http.Request) {
	ClearTable(r.DB)
	fmt.Fprintf(res, "Success!")
}

func (r *App) RedirectTrixel(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tokenID := vars["tokenID"]
	toke, err := strconv.ParseUint(tokenID, 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	trixel := &Trixel{
		TokenID: toke,
	}
	trixel.GetTrixel(r.DB)
	http.Redirect(res, req, trixel.MetadataUrl, http.StatusMovedPermanently)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS pixels (
	hash varchar(64),
	x smallint unsigned,
	y smallint unsigned,
	color varchar(8),
	last_address varchar(64)
)`

func EnsureTableExists(db *gorm.DB) {
	if err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func ClearTable(db *gorm.DB) {
	db.Exec("DELETE FROM pixels")
}
