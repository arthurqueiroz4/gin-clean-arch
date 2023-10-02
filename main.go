package main

import (
	"fmt"
	"gin-clean-arch/bootstrap"
)

func main() {
	app := bootstrap.App()
	bootstrap.NewPostgresDatabase(app.Env)
	bootstrap.GenerateDatabase(app.Database)
	fmt.Println(app.Database)
}
