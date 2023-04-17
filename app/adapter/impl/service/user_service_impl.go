package service

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/repository"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) service.UserService {
	return &userServiceImpl{userRepository: ur}
}

func (s *userServiceImpl) Index() ([]model.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
