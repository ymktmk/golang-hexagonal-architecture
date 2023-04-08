package service

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/repository"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
)

type gettingUserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewGettingUserService(ur repository.UserRepository) service.GettingUserService {
	return &gettingUserServiceImpl{userRepository: ur}
}

func (s *gettingUserServiceImpl) Index() ([]model.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
