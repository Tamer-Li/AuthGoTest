package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func New() *API {
	api := API{
		router: mux.NewRouter(),
	}

	api.router.HandleFunc("/", Hello).Methods(http.MethodGet)

	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "<html><body><h2>Hello World</h2></body></html>")
	if err != nil {
		panic(err)
	}
}
