// Package store contains the functionality for accessing the database
package store

import (
	"articlemaker/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"os"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
)

// uncomment driver import based on your needs

type databaseConfigurations struct {
	dbdriver   string
	dbname     string
	dbuser     string
	dbpassword string
	dbhost     string
	dbport     string
	env        string
}

// SetUpDb creates the schema tables, gorm guarantees the initial migrations run only once
func SetUpDb() {
	db := GetConnection()
	defer db.Close()

	db.AutoMigrate(&models.Publisher{}, &models.Category{}, &models.Article{})
	db.Model(&models.Article{}).AddForeignKey("publisher_id", "publishers(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Article{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
}

// GetConnection get database connection
func GetConnection() *gorm.DB {
	dsn := getDataSourceName()
	configs := getDbConfigurations()
	db, err := gorm.Open(configs.dbdriver, dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	return db
}

// getDataSourceName get the database type
func getDataSourceName() (dsn string) {
	configs := getDbConfigurations()

	switch configs.dbdriver {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			configs.dbuser, configs.dbpassword, configs.dbhost, configs.dbname,
		)
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			configs.dbhost, configs.dbport, configs.dbuser, configs.dbname, configs.dbpassword,
		)
	case "mssql":
		dsn = fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			configs.dbuser, configs.dbpassword, configs.dbhost, configs.dbport, configs.dbname,
		)
	case "sqlite":
		dsn = fmt.Sprintf("/tmp/%s.db", configs.dbname)
	}

	return dsn
}

// getDbConfigurations get the database configurations from environment
func getDbConfigurations() (configs databaseConfigurations) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	configs.dbdriver = viper.GetString(fmt.Sprintf("database.%s.dbdriver", env))
	configs.dbuser = viper.GetString(fmt.Sprintf("database.%s.dbuser", env))
	configs.dbpassword = viper.GetString(fmt.Sprintf("database.%s.dbpassword", env))
	configs.dbhost = viper.GetString(fmt.Sprintf("database.%s.dbhost", env))
	configs.dbport = viper.GetString(fmt.Sprintf("database.%s.dbport", env))
	configs.dbname = viper.GetString(fmt.Sprintf("database.%s.dbname", env))
	configs.env = env

	return configs
}
