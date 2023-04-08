package usecase

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
)

func RequestGetUser(gettingUserService service.GettingUserService) ([]model.User, error) {
	users, err := gettingUserService.Index()
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
