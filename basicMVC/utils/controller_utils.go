package utils

import "github.com/gin-gonic/gin"

func Respond(ctx *gin.Context, status int, body interface{})  {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(status, body)
		return
	}
	ctx.JSON(status, body)
}

func RespondError(ctx *gin.Context, err *ApplicationError) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(err.Status, err)
		return
	}
	ctx.JSON(err.Status, err)
}