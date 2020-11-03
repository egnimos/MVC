package services

import (
	"github.com/egnimos/mvc/basicMVC/domain"
	"github.com/egnimos/mvc/basicMVC/utils"
)

type userServices struct{}

var (
	UserService userServices
)

func (us *userServices) GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
