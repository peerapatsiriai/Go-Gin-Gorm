package userusecase

import "api/entities"

type UserRepository interface {
	SaveUser(user *entities.User) error
	FindOneUser(id int) (*entities.User, error)
	FindAllUser() ([]*entities.User, error)
	UpdateOneUser(user *entities.User) error
	DeleteOneUser(id int) error
}
