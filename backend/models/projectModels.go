package models

import (
	"time"

	"gorm.io/gorm"
)

// Project represents a client or admin owned collaboration space.
type Project struct {
	gorm.Model
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Owner       User   `json:"owner" gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	Status      string `json:"status" gorm:"default:ACTIVE"`
	IsPublic    bool   `json:"is_public" gorm:"default:false"`
	CoverImage  string `json:"cover_image"`
	Summary     string `json:"summary" gorm:"type:text"`
	ViewCount   int    `json:"view_count" gorm:"default:0"`
	Members     []ProjectMember      `json:"members" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Invitations []ProjectInvitation  `json:"invitations" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}

// ProjectMember links a user to a project with a role.
type ProjectMember struct {
	gorm.Model
	ProjectID uint      `json:"project_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Role      string    `json:"role" gorm:"default:MEMBER"`
	JoinedAt  time.Time `json:"joined_at"`
}

// ProjectInvitation captures invites to projects.
type ProjectInvitation struct {
	gorm.Model
	ProjectID    uint       `json:"project_id" gorm:"not null"`
	InviterID    uint       `json:"inviter_id" gorm:"not null"`
	InviteeID    *uint      `json:"invitee_id"`
	InviteeEmail string     `json:"invitee_email"`
	Role         string     `json:"role" gorm:"default:COLLABORATOR"`
	Status       string     `json:"status" gorm:"default:PENDING"`
	Message      string     `json:"message" gorm:"type:text"`
	ExpiresAt    *time.Time `json:"expires_at"`
	ResendCount  int        `json:"resend_count" gorm:"default:0"`
	RespondedAt  *time.Time `json:"responded_at"`
}

// ProjectActivity captures project events for feeds.
type ProjectActivity struct {
	gorm.Model
	ProjectID uint   `json:"project_id" gorm:"not null"`
	ActorID   *uint  `json:"actor_id"`
	Action    string `json:"action" gorm:"not null"`
	Details   string `json:"details" gorm:"type:text"`
}

// ProjectChat links a project to a chat.
type ProjectChat struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProjectID uint      `json:"project_id" gorm:"not null"`
	ChatID    string    `json:"chat_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// ProjectTask represents a task within a project.
type ProjectTask struct {
	gorm.Model
	ProjectID   uint       `json:"project_id" gorm:"not null"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description" gorm:"type:text"`
	Status      string     `json:"status" gorm:"default:OPEN"`
	Priority    string     `json:"priority" gorm:"default:MEDIUM"`
	AssigneeID  *uint      `json:"assignee_id"`
	DueDate     *time.Time `json:"due_date"`
}

// ProjectDeliverable represents a deliverable milestone.
type ProjectDeliverable struct {
	gorm.Model
	ProjectID   uint       `json:"project_id" gorm:"not null"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description" gorm:"type:text"`
	Status      string     `json:"status" gorm:"default:PENDING"`
	DueDate     *time.Time `json:"due_date"`
}

// ProjectSubmission represents a deliverable submission by a member.
type ProjectSubmission struct {
	gorm.Model
	DeliverableID uint       `json:"deliverable_id" gorm:"not null"`
	ProjectID     uint       `json:"project_id" gorm:"not null"`
	SubmitterID   uint       `json:"submitter_id" gorm:"not null"`
	Content       string     `json:"content" gorm:"type:text"`
	FileURL       string     `json:"file_url"`
	Status        string     `json:"status" gorm:"default:SUBMITTED"`
	ReviewedBy    *uint      `json:"reviewed_by"`
	ReviewedAt    *time.Time `json:"reviewed_at"`
	Feedback      string     `json:"feedback" gorm:"type:text"`
}
