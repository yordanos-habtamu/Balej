package models

import (
	"gorm.io/gorm"
)

// User represents a user who can offer or apply for jobs
type User struct {
	gorm.Model
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	Role        string    `json:"role" gorm:"type:user_role;default:'talent'"`
	CV          string    `json:"cv"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	IsVerified        bool      `json:"is_verified" gorm:"default:false"`
	VerificationToken string    `json:"verification_token"`
	TwoFactorEnabled  bool      `json:"two_factor_enabled" gorm:"default:false"`
	TwoFactorSecret   string    `json:"two_factor_secret"`
	ThemePreference   string    `json:"theme_preference" gorm:"default:'talent_purple'"`
}

