package tests

import (
	"articlemaker/server"
	"articlemaker/store"
	"github.com/spf13/viper"
)

func App() {
	app := &server.Server{
		Host: viper.GetString("server.testing.appurl"),
		Port: viper.GetString("server.testing.port"),
	}

	app.Start()
}

func TearDown() {
	db := store.GetConnection()
	db.DropTable("articles", "categories", "publishers")
}

func SetUp() {
	store.SetUpDb()
}
