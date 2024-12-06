package api

import "github.com/gorilla/mux"

func New() *API {
	api := API{
		router: mux.NewRouter(),
	}

	return &api
}
