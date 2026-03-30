package models

import "time"

// FileAsset represents an uploaded file for chat or projects.
type FileAsset struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	OwnerID   uint       `json:"owner_id" gorm:"not null"`
	ProjectID *uint      `json:"project_id"`
	ChatID    *string    `json:"chat_id"`
	Context   string     `json:"context" gorm:"default:GENERAL"`
	FileName  string     `json:"file_name" gorm:"not null"`
	FileURL   string     `json:"file_url" gorm:"not null"`
	FileType  string     `json:"file_type"`
	FileSize  int64      `json:"file_size"`
	Tags      string     `json:"tags" gorm:"type:text"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

// VideoSession represents a scheduled or live video call.
type VideoSession struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	ProjectID   *uint      `json:"project_id"`
	ChatID      *string    `json:"chat_id"`
	HostID      uint       `json:"host_id" gorm:"not null"`
	Title       string     `json:"title" gorm:"not null"`
	MeetingURL  string     `json:"meeting_url"`
	Status      string     `json:"status" gorm:"default:SCHEDULED"`
	ScheduledAt time.Time  `json:"scheduled_at" gorm:"not null"`
	StartedAt   *time.Time `json:"started_at"`
	EndedAt     *time.Time `json:"ended_at"`
	RecordingURL string    `json:"recording_url"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// VideoSessionParticipant represents attendees for a session.
type VideoSessionParticipant struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SessionID uint      `json:"session_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Role      string    `json:"role" gorm:"default:ATTENDEE"`
	Status    string    `json:"status" gorm:"default:INVITED"`
	JoinedAt  *time.Time `json:"joined_at"`
}

// VideoSessionInvite represents an invitation to a session.
type VideoSessionInvite struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	SessionID  uint       `json:"session_id" gorm:"not null"`
	InviterID  uint       `json:"inviter_id" gorm:"not null"`
	InviteeID  *uint      `json:"invitee_id"`
	InviteeEmail string   `json:"invitee_email"`
	Status     string     `json:"status" gorm:"default:PENDING"`
	Message    string     `json:"message"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	RespondedAt *time.Time `json:"responded_at"`
}
