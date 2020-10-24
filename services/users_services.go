package services

import (
	"github.com/egnimos/mvc/domain"
	"github.com/egnimos/mvc/utils"
)

func GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}