package service

import "github.com/alejandroagarcia/GO-Twitter/src/domain"

type UserManager struct {
	Users []*domain.User
}

func NewUserManager() *UserManager {
	var um *UserManager
	um.Users = make([]*domain.User, 0)
	return um
}

func (um *UserManager) GetUsers() []*domain.User {
	return um.Users
}

func (um *UserManager) RegisterUser(name string, email string, nick string, password string) {
	var user *domain.User

	user.Name = name
	user.Email = email
	user.Nick = nick
	user.Password = password

	um.Users = append(um.Users, user)
}
