package service

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/repository"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/service"
)

type GettingUserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewGettingUserService(ur repository.UserRepository) service.GettingUserService {
	return &GettingUserServiceImpl{userRepository: ur}
}

func (s *GettingUserServiceImpl) Index() ([]model.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
