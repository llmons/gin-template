package main

import (
	"fmt"

	"github.com/llmons/gin-template/internal/base/conf"
)

func main() {
	c, err := conf.ReadConfig("./configs/config.yaml")
	if err != nil {
		panic(err)
	}

	app, cleanup, err := initApplication(c.Debug, &c.Data.Database, c.Server)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	fmt.Println("Starting server at", c.Server.HTTP.Addr)
	if err := app.Run(c.Server.HTTP.Addr); err != nil {
		panic(err)
	}
}
