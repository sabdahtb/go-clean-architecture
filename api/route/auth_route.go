package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"sabda.go/learn/api/controller"
	"sabda.go/learn/bootstrap"
	"sabda.go/learn/repository"
	"sabda.go/learn/usecase"
)

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ar := repository.NewAuthRepository()
	lc := &controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(ar, timeout),
		Env:         env,
	}

	group.GET("/login", lc.Login)
}
