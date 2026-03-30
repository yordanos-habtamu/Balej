package models

import "time"

// FeatureFlag controls feature access.
type FeatureFlag struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Key         string    `json:"key" gorm:"unique;not null"`
	Enabled     bool      `json:"enabled" gorm:"default:false"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// SkillCategory provides admin-managed skill categories.
type SkillCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// CompliancePolicy represents admin-defined policy rules.
type CompliancePolicy struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content" gorm:"type:text"`
	Status      string    `json:"status" gorm:"default:ACTIVE"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// ComplianceLog captures compliance actions.
type ComplianceLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ActorID   *uint     `json:"actor_id"`
	Action    string    `json:"action"`
	Details   string    `json:"details" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
