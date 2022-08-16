package config

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/dally469/api/packages/models"

	
)

var DB *gorm.DB
func ConnectDb() {
	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("database"))
	fmt.Println(connectionUrl)
	database, err := gorm.Open("mysql", connectionUrl)
	if err != nil {
		panic("Database connection error " + err.Error())
	}

	database.Model(&models.Book{}).AddForeignKey("author_id", "authors(id)", "RESTRICT", "RESTRICT")

	database.AutoMigrate(
		&models.Book{},
		&models.Author{},
	)

	

	DB = database
	
}