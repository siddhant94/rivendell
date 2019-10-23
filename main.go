package main

import (
	"fmt"
	"rivendell/config"
	"rivendell/app"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	fmt.Println("Server Running on port" + config.Port)
	app.Run(config.Port)
}