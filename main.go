package main

import (
	"config"
	"app"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(tracer.WithDebugMode(true))
        myConfig := config.GetConfig()

        myApp := &app.App{}
        myApp.Initialize(myConfig)
        myApp.Run(":3000")
        defer tracer.Stop()
}
