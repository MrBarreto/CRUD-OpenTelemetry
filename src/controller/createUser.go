package controller

import (
	"fmt"

	"github.com/MrBarreto/CRUD-OpenTelemetry/src/configuration/validation"
	"github.com/MrBarreto/CRUD-OpenTelemetry/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		Erro := validation.ValidateUserError(err)
		c.JSON(Erro.Code, Erro)
		return
	}
	fmt.Println(UserRequest)

}
