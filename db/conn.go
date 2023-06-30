package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database connection string 
const DSN = "host=localhost user=unicorn_user password=magical_password dbname=rainbow_database port=5050 sslmode=disable"

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//db.AutoMigrate(&models.User{}, &models.Property{})

	Database = DbInstance{Db: db}
}