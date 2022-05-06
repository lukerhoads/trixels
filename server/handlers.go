package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Router struct {
	*Store
	*Contract
}

func NewRouter(store *Store) *Router {
	return &Router {
		Store: store,
	}
}

func (r *Router) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/pixels", r.HandleGetPixels).Methods("GET")
}

func (r *Router) HandleGetPixels(res http.ResponseWriter, req *http.Request) {
	// Fetch pixels from contract
	
	// Apply the commits in MySql
	commits := r.Store.GetDayCommits()

	// Return as JSON list
}