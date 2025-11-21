package router

import (
	"apispotify/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/damso", controller.Damso)
	mux.HandleFunc("/laylow", controller.Laylow)

	return mux
}
