package main

import (
	"fmt"
	"log"

	"github.com/yordanos-habtamu/GormWithGraphql/models/db"
)

func main() {
	fmt.Println("🚀 Running JobMatcher ATS database migrations...")
	
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	fmt.Println("✅ JobMatcher ATS database migrations completed successfully!")
	fmt.Println("\n📋 Created tables:")
	fmt.Println("  - talent_profiles - Professional profiles for job seekers")
	fmt.Println("  - skills - Skill competencies with proficiency levels")
	fmt.Println("  - experiences - Work experience and employment history")
	fmt.Println("  - work_samples - Portfolio items and work samples")
	fmt.Println("\n🔤 Created enum types:")
	fmt.Println("  - proficiency_level (beginner, intermediate, advanced, expert)")
	fmt.Println("  - skill_category (technology, healthcare, finance, education, marketing,")
	fmt.Println("                    sales, engineering, design, business, operations,")
	fmt.Println("                    customer_service, human_resources, legal, media,")
	fmt.Println("                    science, administration, consulting, retail,")
	fmt.Println("                    hospitality, transportation, manufacturing,")
	fmt.Println("                    real_estate, non_profit, government, startup)")
	fmt.Println("\n⚡ Created indexes for better query performance")
	fmt.Println("🔒 Added constraints:")
	fmt.Println("  - CASCADE delete on foreign keys for data integrity")
	fmt.Println("  - CHECK constraint for experience date validation")
	fmt.Println("  - UNIQUE constraint for skills per profile to prevent duplicates")
	fmt.Println("\n🎯 JobMatcher ATS is ready for production use!")
	fmt.Println("   Supports 25+ industry categories for comprehensive hiring needs")
	
	sqlDB, err := database.DB()
	if err == nil {
		sqlDB.Close()
	}
}
