// Package server sets up everything needed to get the server up and running
package server

import (
	"articlemaker/routes"
	"articlemaker/store"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strings"
)

// Server holds API host information
type Server struct {
	Host string
	Port string
}

func init() {
	viper.SetConfigName("config")
	path, patherror := os.Getwd()

	if patherror != nil {
		panic(fmt.Errorf("Fatal getting current directory: %s \n", patherror))
	}
	if strings.Contains(path, "tests") {
		path = "../"
	} else {
		path = "."
	}
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}
	viper.WatchConfig()

	store.SetUpDb()
}

// Start bootstraps the app and gets it running
func (server *Server) Start() {
	router := routes.HandleRequests()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server.Host, server.Port), router))
}
