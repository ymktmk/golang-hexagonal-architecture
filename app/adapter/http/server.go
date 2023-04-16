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
	userService := service.NewUserService(repository.NewUserRepository(db))
	userController := controller.NewUserController(userService)

	Router = mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/users", userController.GetUser).Methods("GET")
}

func Serve() {
	log.Fatal(http.ListenAndServe(":8000", Router))
}
