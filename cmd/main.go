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
	dbs := app.Postgree
	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, gin, dbs)

	gin.Run(env.ServerAddress)
}
