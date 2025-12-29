package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the PostgreSQL database and sets up the schema
func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create enum types if they don't exist
	createEnumIfNotExists(db, "user_role", []string{"admin", "talent", "client"})
	createEnumIfNotExists(db, "job_type", []string{"full_time", "part_time", "contract", "internship"})
	createEnumIfNotExists(db, "job_category", []string{"construction", "repairing", "house_keeping", "management"})
	createEnumIfNotExists(db, "application_status", []string{"pending", "accepted", "rejected"})

	// Auto-migrate tables
	err = db.AutoMigrate(&models.User{}, &models.JobOffer{}, &models.JobApplicant{}, &models.Chat{},
		&models.Message{})
	if err != nil {
		panic("failed to auto-migrate")
	}

	return db, nil
}

// createEnumIfNotExists checks if an enum type exists and creates it if it doesn't
func createEnumIfNotExists(db *gorm.DB, typeName string, values []string) {
	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_type WHERE typname = ?", typeName).Scan(&count)
	if count == 0 {
		// Build the CREATE TYPE query
		valuesStr := "'" + strings.Join(values, "', '") + "'"
		query := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", typeName, valuesStr)
		if err := db.Exec(query).Error; err != nil {
			panic(fmt.Sprintf("failed to create enum type %s: %v", typeName, err))
		}
	}
}
