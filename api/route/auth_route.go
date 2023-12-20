package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sabda.go/learn/api/controller"
	"sabda.go/learn/bootstrap"
	"sabda.go/learn/repository"
	"sabda.go/learn/usecase"
)

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, route *gin.RouterGroup) {
	ar := repository.NewAuthRepository()
	ac := &controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(ar, timeout),
		Env:         env,
		DB:          db,
	}

	route.GET("/login", ac.Login)
}
