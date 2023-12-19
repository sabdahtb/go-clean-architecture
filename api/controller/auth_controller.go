package controller

import (
	"net/http"

	"sabda.go/learn/bootstrap"
	"sabda.go/learn/domain"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
	Env         *bootstrap.Env
}

func (ac *AuthController) Login(c *gin.Context) {

	loginResponse := domain.SuccessResponse{
		Message: "hello world",
	}

	c.JSON(http.StatusOK, loginResponse)
}
