package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sabda.go/learn/bootstrap"
	"sabda.go/learn/domain"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase
	Env         *bootstrap.Env
	DB          *gorm.DB
}

func (ac *AuthController) Login(c *gin.Context) {

	token, err := ac.AuthUsecase.CreateAccessToken(c, ac.DB, &domain.User{Name: "sabdahtb", Email: "sabda@yopmail.id", Password: "password"})

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := &domain.SuccessResponse{
		Message: token,
	}

	c.JSON(http.StatusOK, loginResponse)
}
