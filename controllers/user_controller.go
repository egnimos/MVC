package controllers

import (
	"github.com/egnimos/mvc/services"
	"github.com/egnimos/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//controllers handle the requests from the client
func GetUsers( ctx *gin.Context) {
	userId := ctx.Param("user_id")
	//convert the string into int
	id, err := strconv.Atoi(userId)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "UserId must be a number",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(ctx, apiErr)
		return
	}
	//getting the user of that particular ID
	user, apiErr := services.UserService.GetUser(id)
	if apiErr != nil {
		//send the bad request
		utils.RespondError(ctx, apiErr)
		return
	}
	//print the result
	utils.Respond(ctx, http.StatusCreated, user)
}

