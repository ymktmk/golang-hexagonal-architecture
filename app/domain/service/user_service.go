package service

import "github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"

type UserService interface {
	Index() ([]model.User, error)
}
