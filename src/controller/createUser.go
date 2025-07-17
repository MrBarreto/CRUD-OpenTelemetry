package controller

import (
	"fmt"

	customError "github.com/MrBarreto/CRUD-OpenTelemetry/src/configuration/custom_error"
	"github.com/MrBarreto/CRUD-OpenTelemetry/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		Erro := customError.NewBadRequestError(
			fmt.Sprintf("Campos inválidos, %s\n", err.Error()))
		c.JSON(Erro.Code, Erro)
		return
	}
	fmt.Println(UserRequest)

}
