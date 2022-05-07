package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Router struct {
	*Store
	*Trixels
}

func NewRouter(store *Store) *Router {
	return &Router {
		Store: store,
	}
}

func (r *Router) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/pixels", r.HandleGetPixels).Methods("GET")
	router.HandleFunc("/commit", r.CreateCommit).Methods("POST")
	router.HandleFunc("/trigger/mint", r.HandleTriggerMint).Methods("GET")
	router.HandleFunc("/trigger/update", r.HandleTriggerUpdate).Methods("GET")
}

func (r *Router) HandleGetPixels(res http.ResponseWriter, req *http.Request) {
	// Fetch pixels from contract
	pixels, err := r.Trixels.GetPixels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Apply the commits in MySql
	commits := r.Store.GetDayCommits()
	for _, c := range commits {
		pixels[c.X][c.Y] = c.Color
	}

	// Return as JSON list
	var dimensionedPixels []DimensionedPixel 
	int index = 0;
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			dimensionedPixels[index] = DimensionedPixel {
				X: string(j),
				Y: string(i),
				Color: pixels[i][j].Color,
			}
		}
	}

	json.NewEncoder(w).Encode(dimensionedPixels)
	return
}

func (r *Router) CreateCommit(res http.ResponseWriter, req *http.Request) {
	var c Commit 
	err := json.NewDecoder(r.Body).Decode(&c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if err := r.Store.CreateCommit(c.X, c.Y, c.Color, c.Address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
}

func (r *Router) HandleTriggerMint(res http.ResponseWriter, req *http.Request) {
	r.devChan<-"mint"
}

func (r *Router) HandleTriggerUpdate(res http.ResponseWriter, req *http.Request) {
	r.devChan<-"updatepixels"
}