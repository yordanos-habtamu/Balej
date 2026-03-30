package models

import (
	"time"

	"gorm.io/gorm"
)

// TalentProfile represents a talent's professional profile
type TalentProfile struct {
	gorm.Model
	UserID            uint   `json:"user_id" gorm:"unique;not null"`
	User              User   `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	ProfessionalTitle string `json:"professional_title"`
	Bio               string `json:"bio" gorm:"type:text"`
	Location          string `json:"location"`
	ProfileImageURL   string `json:"profile_image_url"`
	IsPublic          bool   `json:"is_public" gorm:"default:false"`
	ViewCount         int    `json:"view_count" gorm:"default:0"`
	CompletionPercent int    `json:"completion_percent" gorm:"default:0"`
	Skills            []Skill      `json:"skills" gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE"`
	Experiences       []Experience `json:"experiences" gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE"`
	WorkSamples       []WorkSample `json:"work_samples" gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE"`
}

// Skill represents a specific competency or expertise area
type Skill struct {
	gorm.Model
	ProfileID        uint   `json:"profile_id" gorm:"not null"`
	SkillName        string `json:"skill_name" gorm:"not null"`
	ProficiencyLevel string `json:"proficiency_level" gorm:"type:proficiency_level"` // Beginner, Intermediate, Advanced, Expert
	Category         string `json:"category" gorm:"type:skill_category"`              // Construction, House Keeping, Repairing, Management, Other
}

// Experience represents a record of previous employment or project work
type Experience struct {
	gorm.Model
	ProfileID   uint       `json:"profile_id" gorm:"not null"`
	JobTitle    string     `json:"job_title" gorm:"not null"`
	CompanyName string     `json:"company_name" gorm:"not null"`
	StartDate   time.Time  `json:"start_date" gorm:"not null"`
	EndDate     *time.Time `json:"end_date"` // Nullable for current positions
	Description string     `json:"description" gorm:"type:text"`
	IsCurrent   bool       `json:"is_current" gorm:"default:false"`
}

// WorkSample represents a portfolio item showcasing talent's previous work
type WorkSample struct {
	gorm.Model
	ProfileID    uint   `json:"profile_id" gorm:"not null"`
	Title        string `json:"title"`
	Description  string `json:"description" gorm:"type:text"`
	ImageURL     string `json:"image_url" gorm:"not null"`
	ExternalLink string `json:"external_link"`
	DisplayOrder int    `json:"display_order" gorm:"default:0"`
}
