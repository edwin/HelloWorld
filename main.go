package main

import (
	"config"
	"app"
)

func main() {
	myConfig := config.GetConfig()

	myApp := &app.App{}
	myApp.Initialize(myConfig)
	myApp.Run(":3000")

}