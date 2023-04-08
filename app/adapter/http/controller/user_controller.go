package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/db"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/http/view"
	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/impl/repository"
	serviceImpl "github.com/ymktmk/golang-hexagonal-architecture/app/adapter/impl/service"
	"github.com/ymktmk/golang-hexagonal-architecture/app/application/usecase"
)

func ControlGetUser(w http.ResponseWriter, r *http.Request) {

	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	gettingUserService := serviceImpl.NewGettingUserService(repository.NewUserRepository(db))
	users, err := usecase.RequestGetUser(gettingUserService)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	v := view.NewUser(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
