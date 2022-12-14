package database

import (
	"github.com/woohaen88/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	var DSN string = "root:password@tcp(127.0.0.1:3306)/app?parseTime=true"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Todo: Add migrations
	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal(err)
	}
	Database = DBInstance{DB: db}

}
