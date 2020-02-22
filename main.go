package main

import (
	"articlemaker/server"
	"github.com/spf13/viper"
)

func main() {
	app := &server.Server{
		Host: viper.GetString("server.appurl"),
		Port: viper.GetString("server.port"),
	}

	app.Start()
}
