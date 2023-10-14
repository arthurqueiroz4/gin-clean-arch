package main

import (
	"gin-clean-arch/api/route"
	"gin-clean-arch/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	gin := gin.Default()

	route.Setup(env, app.Database, gin)

	gin.Run(env.ServerAddress)
}
