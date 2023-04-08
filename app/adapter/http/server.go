package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/http/controller"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/users", controller.ControlGetUser).Methods("GET")
}

func Serve() {
	log.Fatal(http.ListenAndServe(":8000", Router))
}
