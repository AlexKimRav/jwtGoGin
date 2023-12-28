package repository

import "jwtgogin/model"

type UserRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindAll() []model.Users
	FindById(usersId int) (model.Users, error)
	FindByUsername(username string) (model.Users, error)
}
