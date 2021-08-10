package config

import (
	"alta/project/model"
	_ "alta/project/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {

	// Framecode for environmental
	/*
		envVar := "root:Minus12345@tcp(localhost:3306)/new_schema?charset=utf8&parseTime=True&loc=Local"
		connectionString := os.Getenv(envVar)
	*/

	//Set connection string here, use mysql username password and schema at your pc
	//connectionString := "root:Minus12345@tcp(localhost:3306)/alta_shop_project?charset=utf8&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")) // doni local computer
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Shopping_cart{})
}
