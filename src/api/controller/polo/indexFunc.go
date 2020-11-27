package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type welcome struct {
	Message string
	CheckTest string
}

//var (
//	lal = []&welcome{}
//)

func IndexFunction(ctx *gin.Context) {
	w := &welcome {
		Message: "Welcome to the microservices practice",
		CheckTest: "/createRepo",
	}
	ctx.JSON(http.StatusOK, w)
}