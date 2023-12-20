package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sabda.go/learn/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine, db *gorm.DB) {
	authRouter := gin.Group("v1/auth")
	NewAuthRouter(env, timeout, db, authRouter)

	productrouter := gin.Group("v1/product")
	NewProductRouter(env, timeout, db, productrouter)
}
