package view

import "github.com/ymktmk/golang-hexagonal-architecture/app/domain/model"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func NewUser(users []model.User) []User {
	res := make([]User, 0, len(users))
	for _, v := range users {
		user := User{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
		res = append(res, user)
	}
	return res
}
