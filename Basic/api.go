package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	r mux.Router
}

func (api *Api) init() {
	api.r = mux.NewRouter()
	http.Handle("/", r)
}
