package db

import (
	"fmt"
	"os"
	"strings"
    "github.com/joho/godotenv"
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)



// ConnectDB establishes a connection to the PostgreSQL database and sets up the schema
func ConnectDB() (*gorm.DB, error) {
	fmt.Println("🔗 Connecting to JobMatcher ATS database...")

	err := godotenv.Load()
	if err != nil {
	
		err = godotenv.Load("../../.env")
	}
	if err != nil {
		// Try loading from ../.env
		err = godotenv.Load("../.env")
	}

	// If still error, check if critical env vars are set before failing
	if err != nil && os.Getenv("DB_HOST") == "" {
		log.Fatal("Error loading .env file: ", err)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
    fmt.Printf("📍 Database: %s@%s:%s/%s\n", user, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to JobMatcher ATS database")
	}

	fmt.Println("✅ Database connection established successfully!")
	fmt.Println("🏗️  Setting up JobMatcher ATS schema...")

	// Create enum types if they don't exist
	createEnumIfNotExists(db, "user_role", []string{"admin", "talent", "client"})
	createEnumIfNotExists(db, "job_type", []string{"full_time", "part_time", "contract", "internship"})
	createEnumIfNotExists(db, "job_category", []string{
		"technology", "healthcare", "finance", "education", "marketing", 
		"sales", "engineering", "design", "business", "operations",
		"customer_service", "human_resources", "legal", "media", "science",
		"administration", "consulting", "retail", "hospitality", "transportation",
		"manufacturing", "real_estate", "non_profit", "government", "startup",
	})
	createEnumIfNotExists(db, "application_status", []string{"pending", "accepted", "rejected"})
	createEnumIfNotExists(db, "job_status", []string{"LIVE", "ARCHIVE", "PENDING"})
	createEnumIfNotExists(db, "notification_type", []string{"TRANSACTIONAL", "BROADCAST", "NEWSLETTER"})
	createEnumIfNotExists(db, "notification_priority", []string{"LOW", "NORMAL", "CRITICAL"})
	createEnumIfNotExists(db, "proficiency_level", []string{"beginner", "intermediate", "advanced", "expert"})
	createEnumIfNotExists(db, "skill_category", []string{
		"technology", "healthcare", "finance", "education", "marketing", 
		"sales", "engineering", "design", "business", "operations",
		"customer_service", "human_resources", "legal", "media", "science",
		"administration", "consulting", "retail", "hospitality", "transportation",
		"manufacturing", "real_estate", "non_profit", "government", "startup",
	})

	// Auto-migrate tables
	err = db.AutoMigrate(&models.User{}, &models.JobOffer{}, &models.JobApplicant{}, &models.Chat{},
		&models.Message{}, &models.TalentProfile{}, &models.Skill{}, &models.Experience{}, &models.WorkSample{},
		&models.Notification{}, &models.NewsletterSubscription{}, &models.Flag{}, &models.AuditLog{})
	if err != nil {
		panic("❌ Failed to auto-migrate JobMatcher ATS tables")
	}

	// Create indexes for profile-related tables
	createIndexes(db)

	fmt.Println("🎯 JobMatcher ATS schema setup complete!")
	fmt.Println("📊 Ready to handle 25+ industry categories and comprehensive hiring workflows")
	
	return db, nil
}

// createIndexes creates database indexes for better query performance
func createIndexes(db *gorm.DB) {
	// Index for talent_profiles
	db.Exec("CREATE INDEX IF NOT EXISTS idx_talent_profiles_user_id ON talent_profiles(user_id)")
	
	// Indexes for skills
	db.Exec("CREATE INDEX IF NOT EXISTS idx_skills_profile_id ON skills(profile_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_skills_category ON skills(category)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_skills_skill_name ON skills(skill_name)")
	
	// Indexes for experiences
	db.Exec("CREATE INDEX IF NOT EXISTS idx_experiences_profile_id ON experiences(profile_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_experiences_start_date ON experiences(start_date)")
	
	// Indexes for work_samples
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_samples_profile_id ON work_samples(profile_id)")
	
	// Add CHECK constraint for experience date validation
	db.Exec("ALTER TABLE experiences DROP CONSTRAINT IF EXISTS check_experience_dates")
	db.Exec("ALTER TABLE experiences ADD CONSTRAINT check_experience_dates CHECK (end_date IS NULL OR end_date >= start_date)")
	
	// Add unique constraint for skills (profile_id, skill_name combination)
	db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_skills_profile_skill_unique ON skills(profile_id, skill_name) WHERE deleted_at IS NULL")
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
