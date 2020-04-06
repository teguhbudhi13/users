package main

import (
	"github.com/teguhbudhi13/users/app"
	"github.com/teguhbudhi13/users/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + config.Server.Port)
}
