package config

import (
	"final-project/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := "root"
	password := "1d1--63-2gEgfCbDH2adaFeagaG1-b3H"
	host := "tcp(viaduct.proxy.rlwy.net:52716)"
	database := "railway"

	// username := "root"
	// password := ""
	// host := "tcp(127.0.0.1:3306)"
	// database := "restaurant_db"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Restaurant{}, &models.Review{}, &models.User{}, &models.Menu{}, &models.OrderHistory{})

	return db
}
