package store

import (
	"articlemaker/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
     _ "github.com/go-sql-driver/mysql"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
)
// uncomment driver import based on your needs


type databaseConfigurations struct {
	dbdriver string
	dbname string
	dbuser string
	dbpassword string
	dbhost string
	dbport string
}

// SetUpDb creates the schema tables, gorm guarantees the initial migrations run only once
func SetUpDb() {
	db := GetConnection()
	defer db.Close()

	db.AutoMigrate(&models.Publisher{}, &models.Category{}, &models.Article{})
	//db.Model(&models.Article{}).AddForeignKey("publisher_id", "publishers(id)", "ALLOW", "RESTRICT")
	//db.Model(&models.Article{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
}

func GetConnection() *gorm.DB {
	dsn := getDataSourceName()
	db, err := gorm.Open(viper.GetString("database.dbdriver"), dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	return db
}

func getDataSourceName() (dsn string) {
	configs := getDbConfigurations()

	switch configs.dbdriver {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
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
			configs.dbuser, configs.dbpassword, configs.dbhost, configs.dbport,configs. dbname,
		)
	case "sqlite":
		dsn = fmt.Sprintf("/tmp/%s.db", configs.dbname)
	}

	return dsn
}

func getDbConfigurations() (configs databaseConfigurations) {
	configs.dbdriver = viper.GetString("database.dbdriver")
	configs.dbname = viper.GetString("database.dbname")
	configs.dbuser = viper.GetString("database.dbuser")
	configs.dbpassword = viper.GetString("database.dbpassword")
	configs.dbhost = fmt.Sprintf("(%s)", viper.GetString("database.dbhost"))
	configs.dbport = viper.GetString("database.dbport")

	return configs
}
