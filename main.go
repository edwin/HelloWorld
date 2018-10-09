package main

import (
	"config"
	"app"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"  
)

func main() {
	tracer.Start(tracer.WithDebugMode(true))

        fmt.Println("1111 1")

        myConfig := config.GetConfig()

        myApp := &app.App{}
        myApp.Initialize(myConfig)
        myApp.Run(":3000")
        fmt.Println("22222 2")
        defer tracer.Stop()
}
