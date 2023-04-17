package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ymktmk/golang-hexagonal-architecture/app/adapter/http/view"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
	"github.com/ymktmk/golang-hexagonal-architecture/app/application/usecase"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := usecase.RequestGetUser(uc.userService)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	v := view.NewUser(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
