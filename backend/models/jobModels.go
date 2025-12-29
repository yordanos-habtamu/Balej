package models
import (
	"time"

	"gorm.io/gorm"
)

// JobApplication represents a job posting
type JobOffer struct {
	gorm.Model
	Type             string    `json:"job_type" gorm:"not null;type:job_type"`
	Title            string    `json:"title" gorm:"not null"`
	Description      string    `json:"description" gorm:"not null;type:text"`
	Location         string    `json:"location" gorm:"not null"`
	OfferedBy        uint      `json:"offered_by" gorm:"not null"`
	OfferedByUser    User      `json:"offered_by_user" gorm:"foreignKey:OfferedBy;references:ID;constraint:OnDelete:RESTRICT"`
	Category         string    `json:"category" gorm:"not null;type:job_category"`
	ApplicationCount uint      `json:"application_count" gorm:"not null;default:0"`
	Applicants     []JobApplicant `json:"applications" gorm:"foreignKey:JobOfferID;constraint:OnDelete:CASCADE"`
}

// JobApplicant represents the join table for the many-to-many relationship between JobApplication and User
type JobApplicant struct {
	JobOfferID uint      `json:"job_id" gorm:"primaryKey;not null"`
	UserID           uint      `json:"user_id" gorm:"primaryKey;not null"`
	AppliedAt        time.Time `json:"applied_at" gorm:"autoCreateTime"`
	Status           string    `json:"status" gorm:"type:application_status;default:'pending'"`
	Resume           string    `json:"resume" gorm:"not null"` 
}

type JobFilter struct {
	JobOfferID uint   `json:"job_offer_id"`
	UserID     uint   `json:"user_id"`
	Category   string `json:"category"`
}
 
type JobApplicantFilter struct {
	JobOfferID uint   `json:"job_offer_id"`
	UserID     uint   `json:"user_id"`
	Status     string `json:"status"`
}