package services

import (
	"github.com/egnimos/mvc/domain"
	"github.com/egnimos/mvc/utils"
)

type userServices struct {}

var (
	UserService userServices
)

func (us *userServices) GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}