package models

import "gorm.io/gorm"

// Role represents a set of permissions.
type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Summary     string `json:"summary" gorm:"type:text"`
	IsSystem    bool   `json:"is_system" gorm:"default:false"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

// Permission represents an access capability.
type Permission struct {
	gorm.Model
	Key         string `json:"key" gorm:"unique;not null"`
	Description string `json:"description" gorm:"type:text"`
}

// RolePermission is join table.
type RolePermission struct {
	RoleID       uint `json:"role_id" gorm:"primaryKey"`
	PermissionID uint `json:"permission_id" gorm:"primaryKey"`
}
