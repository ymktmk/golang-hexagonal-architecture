package service

import "github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"

type GettingUserService interface {
	Index() ([]model.User, error)
}
