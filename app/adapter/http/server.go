package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/db"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/impl/repository"
      "github.com/ymktmk/golang-hexagonal-architecture/app/adapter/impl/service"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/http/controller"
)

var Router *mux.Router

func init() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	gettingUserService := service.NewGettingUserService(repository.NewUserRepository(db))
	controller := controller.NewUserController(gettingUserService)

	Router = mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/users", controller.GetUser).Methods("GET")
}

func Serve() {
	log.Fatal(http.ListenAndServe(":8000", Router))
}
