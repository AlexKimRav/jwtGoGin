package config

import (
	"fmt"
	"jwtgogin/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb(config *Config) *gorm.DB {
	//Getenv is taking from system var with name MYSQL_PASSWORD MYSQL_LOGIN

	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("connected to DB")
	return db
}
