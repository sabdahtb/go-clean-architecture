package main

import (
	"time"

	"github.com/gin-gonic/gin"

	route "sabda.go/learn/api/route"
	"sabda.go/learn/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
