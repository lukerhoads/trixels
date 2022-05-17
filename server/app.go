package main

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

func (a *App) Initialize(db *gorm.DB) {
	a.DB = db
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (r *App) initializeRoutes() {
	r.Router.Use(Middleware)
	r.Router.HandleFunc("/pixels", r.GetPixels).Methods("GET")
	r.Router.HandleFunc("/pixels", r.UpdatePixel).Methods("POST")
	r.Router.HandleFunc("/commit", r.CreateCommit).Methods("POST")
	r.Router.HandleFunc("/trigger/mint", r.TriggerMint).Methods("GET")
	r.Router.HandleFunc("/trigger/update", r.TriggerUpdate).Methods("GET")
	r.Router.HandleFunc("/trigger/reset", r.ResetDB).Methods("GET")
}

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        h.ServeHTTP(w, r)
    })
}

func (a *App) Run(addr string) {
	log.Println("Listening...")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))
}

func (r *App) GetPixels(res http.ResponseWriter, req *http.Request) {
	var pixels []Pixel
	r.DB.Find(&pixels)
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

func (r *App) CreateCommit(res http.ResponseWriter, req *http.Request) {
	var c Commit 
	err := json.NewDecoder(req.Body).Decode(&c)
    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        return
    }

	if err := c.CreateCommit(r.DB); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
        return
	}

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
	log.Println("resetting")
	clearTable(r.DB)
	r.InitializePixels()
	return
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS pixels (
	id bigint unsigned,
	x smallint unsigned,
	y smallint unsigned,
	color varchar(8),
	last_address varchar(64)
)`

func ensureTableExists(db *gorm.DB) {
	if err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable(db *gorm.DB) {
	db.Exec("DELETE FROM pixels")
}

func (a *App) InitializePixels() {
	// ensureTableExists(a.DB)

	for i := 0; i < 30; i++ {
		var pixels []Pixel
		for j := 0; j < 30; j++ {
			pixel := Pixel {
				X: uint16(i),
				Y: uint16(j),
				Color: "#000",
				LastAddress: "0x0",
			}

			pixels = append(pixels, pixel)
		}

		a.DB.Create(&pixels)
	}
}

// func (r *Router) HandleGetPixels(res http.ResponseWriter, req *http.Request) {
// 	// Fetch pixels from contract
// 	colors, err := r.Trixels.Colors()
// 	if err != nil {
// 		http.Error(res, err.Error(), http.StatusBadRequest)
//         return
// 	}

// 	// Apply the commits in MySql
// 	commits := r.Store.GetDayCommits()
// 	for _, c := range commits {
// 		colors[c.X][c.Y] = c.Color
// 	}

// 	// Return as JSON list
// 	var dimensionedPixels []DimensionedPixel 
// 	index := 0
// 	for i := 0; i < len(colors); i++ {
// 		for j := 0; j < len(colors[0]); j++ {
// 			dimensionedPixels[index] = DimensionedPixel {
// 				X: string(j),
// 				Y: string(i),
// 				Color: colors[i][j].Color,
// 			}
// 		}
// 	}

// 	json.NewEncoder(res).Encode(dimensionedPixels)
// 	return
// }
