package user

import "github.com/IsraelTeo/ecommerce-backend/model"

type UseCase interface {
	Create(model *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Storage interface {
	Create(model *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
