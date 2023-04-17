package usecase

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
)

func RequestGetUser(userService service.UserService) ([]model.User, error) {
	users, err := userService.Index()
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
