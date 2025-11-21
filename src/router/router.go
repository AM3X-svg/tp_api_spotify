package router

import (
	"net/http"
	"tpspotify/controller"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	return mux
}
