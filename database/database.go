package database

import (
	"log"

	"github.com/manoelduran/refinancy-with-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connection successfully opened")

    DB.AutoMigrate(&models.Recipe{})
    log.Println("Database migrated")
}

func CloseDatabase() {
	sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("Failed to get database instance:", err)
    }
    err = sqlDB.Close()
    if err != nil {
        log.Println("Failed to close database connection:", err)
    } else {
        log.Println("Database connection successfully closed")
    }
}
