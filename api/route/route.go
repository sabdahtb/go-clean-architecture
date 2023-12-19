package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"sabda.go/learn/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewAuthRouter(env, timeout, publicRouter)
}
