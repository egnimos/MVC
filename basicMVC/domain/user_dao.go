package domain

import (
	"fmt"
	"github.com/egnimos/mvc/utils"
	"net/http"
)

var users = map[int]*User{
	123: {Id: 123, FirstName: "Niteesh", LastName: "Dubey", Email: "niteeshdubey97@gmail.com"},
}

type userDao struct {}

func init() {
	UserDao = &userDao{}
}

var (
	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(userId int) (*User, *utils.ApplicationError)
}

func (ud *userDao) GetUser(userId int) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprint("USER NOT FOUND!!", userId),
		Status: http.StatusNotFound,
		Code: "Not Found",
	}
}
