package repository

import (
	"log"

	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"
	"github.com/ymktmk/golang-hexagonal-architecture/app/domain/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		log.Println(err)
		return []model.User{}, err
	}
	return users, nil
}
