package repository

import "github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
}
