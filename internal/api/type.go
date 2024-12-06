package api

import "github.com/gorilla/mux"

type API struct {
	router *mux.Router
}

func New() *API {
	api := API{
		router: mux.NewRouter(),
	}

	return &api
}
