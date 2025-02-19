package database

import (
	"log"
	"os"

	"github.com/abdullahalsazib/e-com-auth-golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Error in database connection.")
	}
	log.Println("Databse conncection successful.")

	db.AutoMigrate(new(model.User))

	DBConn = db
}
