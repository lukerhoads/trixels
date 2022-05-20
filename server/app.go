package server

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gorm.io/gorm"
	"github.com/gorilla/handlers"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
	devChan chan string
}

func (a *App) Initialize(db *gorm.DB, devChan chan string) {
	a.devChan = devChan
	a.DB = db
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (r *App) initializeRoutes() {
	r.Router.HandleFunc("/pixel/{coord}", r.GetPixel).Methods("GET")
	r.Router.HandleFunc("/pixels", r.GetPixels).Methods("GET")
	r.Router.HandleFunc("/pixels", r.UpdatePixel).Methods("POST")
	r.Router.HandleFunc("/trigger/mint", r.TriggerMint).Methods("GET")
	r.Router.HandleFunc("/trigger/update", r.TriggerUpdate).Methods("GET")
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
	pixel := Pixel {
		X: xCoord,
		Y: yCoord,
	}
	pixel.GetPixel(r.DB)
	json.NewEncoder(res).Encode(pixel)
	return
}

func (r *App) GetPixels(res http.ResponseWriter, req *http.Request) {
	var pixels Pixels
	pixels.GetPixels(r.DB)
	json.NewEncoder(res).Encode(pixels)
	return
}

func (r *App) UpdatePixel(res http.ResponseWriter, req *http.Request) {
	var pixel Pixel 
	err := json.NewDecoder(req.Body).Decode(&pixel)
    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        return
    }

	pixel.UpdatePixel(r.DB)
	fmt.Fprintf(res, "Success!") 
	return
}

func (r *App) TriggerMint(res http.ResponseWriter, req *http.Request) {
	r.devChan<-"mint"
	return
}

func (r *App) TriggerUpdate(res http.ResponseWriter, req *http.Request) {
	r.devChan<-"updatepixels"
	return
}

func (r *App) ResetDB(res http.ResponseWriter, req *http.Request) {
	ClearTable(r.DB)
	r.InitializePixels()
	return
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

func (a *App) InitializePixels() {
	EnsureTableExists(a.DB)

	for i := 0; i < 30; i++ {
		var pixels []Pixel
		for j := 0; j < 30; j++ {
			pixels = append(pixels, *NewPixel(uint16(i), uint16(j)))
		}

		a.DB.Create(&pixels)
	}
}
