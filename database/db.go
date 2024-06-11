package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	OperationalCalendar *gorm.DB
}

var DB DbInstance

func ConnectDb() {

	// Connect to database
	dsn := "host=localhost user=postgres password=Alejandro98 database=operationalcalendar port=5432"

	// Connecting to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger
		Logger: logger.Default.LogMode(logger.Info),
	})

	// If the connection have error
	if err != nil {
		log.Fatal("Failed to connect to database", err)
		os.Exit(2)
	}

	// Create extension UUID
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Connected to db logger
	log.Println("Connected to database")

	// Logger for all info
	db.Logger = logger.Default.LogMode(logger.Info)

	// Logger runing migrations
	log.Println("Running migrations")

	// Migrations
	db.AutoMigrate()

	// Instance of db
	DB = DbInstance{
		OperationalCalendar: db,
	}
}
