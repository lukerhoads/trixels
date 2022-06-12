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
	a.Router = mux.NewRouter()
	a.DB = db
	a.devChan = devChan
	a.initializeRoutes()
}

func (r *App) initializeRoutes() {
	r.Router.HandleFunc("/trixel/live", r.GetLiveTrixel).Methods("GET")
	r.Router.HandleFunc("/trixels", r.GetTrixels).Methods("GET")
	r.Router.HandleFunc("/trixels/find/{trixelID}", r.GetTrixel).Methods("GET")

	r.Router.HandleFunc("/trixel/{trixelID}", r.RedirectTrixel).Methods("GET")
	r.Router.HandleFunc("/pixel/{coord}", r.GetPixel).Methods("GET")
	r.Router.HandleFunc("/pixels", r.GetPixels).Methods("GET")
	r.Router.HandleFunc("/pixels", r.UpdatePixel).Methods("POST")
	r.Router.HandleFunc("/trigger/mint", r.TriggerMint).Methods("GET")
	r.Router.HandleFunc("/trigger/end", r.TriggerEnd).Methods("GET")
	r.Router.HandleFunc("/trigger/reset", r.ResetDB).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Println("Listening on ", addr)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))
}

func (r *App) GetTrixels(res http.ResponseWriter, req *http.Request) {
	var trixels Trixels
	trixels.GetTrixels(r.DB)
	json.NewEncoder(res).Encode(trixels)
}

func (r *App) GetTrixel(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tokenID := vars["trixelID"]
	toke, err := strconv.ParseUint(tokenID, 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	trixel := Trixel{
		TokenID: uint(toke),
	}
	found := trixel.GetTrixel(r.DB)
	if found {
		json.NewEncoder(res).Encode(trixel)
	} else {
		json.NewEncoder(res).Encode(ServerError{
			Message: "Trixel not found",
		})
	}
}

func (r *App) GetLiveTrixel(res http.ResponseWriter, req *http.Request) {
	trixel := Trixel{}
	found := trixel.GetLiveTrixel(r.DB)
	if !found {
		http.Error(res, "Trixel not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(trixel)
}

func (r *App) RedirectTrixel(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tokenID := vars["trixelID"]
	toke, err := strconv.ParseUint(tokenID, 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	trixel := &Trixel{
		TokenID: uint(toke),
	}
	trixel.GetTrixel(r.DB)
	http.Redirect(res, req, trixel.MetadataUrl, http.StatusMovedPermanently)
}

func (r *App) GetPixel(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	coord := vars["coord"]
	xCoord, yCoord := DecodePixelValues(coord)
	pixel := Pixel{
		X: xCoord,
		Y: yCoord,
	}
	found := pixel.GetPixel(r.DB)
	if found {
		json.NewEncoder(res).Encode(pixel)
	} else {
		json.NewEncoder(res).Encode(Pixel{
			X: xCoord,
			Y: yCoord,
			Color: "#FFFFFF",
			Editor: "",
			UpdatedAt: nil,
		})
	}
}

func (r *App) GetPixels(res http.ResponseWriter, req *http.Request) {
	var pixels Pixels
	pixels.GetPixels(r.DB)
	json.NewEncoder(res).Encode(pixels)
}

func (r *App) UpdatePixel(res http.ResponseWriter, req *http.Request) {
	var pixel Pixel
	err := json.NewDecoder(req.Body).Decode(&pixel)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if pixel.X > IMAGE_DIMENSIONS || pixel.Y > IMAGE_DIMENSIONS {
		json.NewEncoder(res).Encode(ServerError{
			Message: "Invalid coordinates",
		})
	}

	matched := HexVerificationRegexp.MatchString(pixel.Color)
	if !matched {
		json.NewEncoder(res).Encode(ServerError{
			Message: "Invalid hex code",
		})
		return
	}

	// Get other pixels that the sender has edited and check if the most recent update time is
	// at least updatetimeinterval ago
	mostRecentPixel := Pixel {
		Editor: pixel.Editor,
	}
	mostRecent := mostRecentPixel.GetMostRecentEdited(r.DB)
	validUpdate := mostRecent.UpdatedAt.Add(PixelUpdateTime).Before(time.Now())
	if !validUpdate {
		json.NewEncoder(res).Encode(ServerError{
			Message: "Cannot update, user has updated another pixel in the last 5 minutes",
		})
		return
	}

	getPixel := Pixel{
		X: pixel.X,
		Y: pixel.Y,
	}
	found := getPixel.GetPixel(r.DB)
	if !found {
		pixel.CreatePixel(r.DB)
	} else {
		validUpdate := getPixel.UpdatedAt.Add(PixelUpdateTime).Before(time.Now())
		if !validUpdate {
			json.NewEncoder(res).Encode(ServerError{
				Message: "Cannot update, pixel has been updated in the last 5 minutes",
			})
			return
		}
	
		pixel.UpdateExistingPixel(r.DB, getPixel.validPixel())
	}

	fmt.Fprintf(res, "Success!")
}

func (r *App) TriggerMint(res http.ResponseWriter, req *http.Request) {
	r.devChan <- "mint"
	fmt.Fprintf(res, "Success!")
}

func (r *App) TriggerEnd(res http.ResponseWriter, req *http.Request) {
	r.devChan <- "end"
	fmt.Fprintf(res, "Success!")
}

func (r *App) ResetDB(res http.ResponseWriter, req *http.Request) {
	ClearTable(r.DB)
	fmt.Fprintf(res, "Success!")
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
	db.Exec("DELETE FROM trixels")
}
