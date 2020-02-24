// Package main is the entry point for the application
package main

import (
	"articlemaker/server"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	app := &server.Server{
		Host: viper.GetString(fmt.Sprintf("server.%s.appurl", env)),
		Port: viper.GetString(fmt.Sprintf("server.%s.port", env)),
	}

	app.Start()
}
