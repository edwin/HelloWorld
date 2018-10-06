package main

import (
	"config"
	"app"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"  
)

func main() {
	myConfig := config.GetConfig()

	myApp := &app.App{}
	myApp.Initialize(myConfig)
	myApp.Run(":3000")

	tracer.Start(tracer.WithServiceName("my-service"))
	defer tracer.Stop()
}
