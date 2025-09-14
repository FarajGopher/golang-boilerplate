package database

import (
	"fmt"
	"golang-boilerplate/config"
	"golang-boilerplate/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB(){
	// Update these with your env variables
	host := config.EnvConfigs.DBHost
	port := config.EnvConfigs.DBPort
	user := config.EnvConfigs.DBUser
	password := config.EnvConfigs.DBPassword
	dbname := config.EnvConfigs.DBName

	//use Spritf to make connection string for database connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate tables
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	DB = db
	fmt.Println("Database connection established and migrated")
	
}