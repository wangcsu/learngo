package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter(api *mux.Router) {
	api.PathPrefix("/raster_data/download/").
		Handler(http.StripPrefix("/raster_data/download/",
			http.FileServer(http.Dir("/Users/frankwang/Desktop/saved/thumbs"))))
}
