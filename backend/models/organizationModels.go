package models

import "time"

// Organization represents a client organization.
type Organization struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OwnerID     uint      `json:"owner_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Industry    string    `json:"industry"`
	Size        string    `json:"size"`
	Website     string    `json:"website"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// OrganizationMember links a user to an organization.
type OrganizationMember struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	Role           string    `json:"role" gorm:"default:MEMBER"`
	JoinedAt       time.Time `json:"joined_at"`
}

// OrganizationInvite represents an invitation.
type OrganizationInvite struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	OrganizationID uint       `json:"organization_id" gorm:"not null"`
	InviterID      uint       `json:"inviter_id" gorm:"not null"`
	InviteeID      *uint      `json:"invitee_id"`
	InviteeEmail   string     `json:"invitee_email"`
	Status         string     `json:"status" gorm:"default:PENDING"`
	Message        string     `json:"message"`
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime"`
	RespondedAt    *time.Time `json:"responded_at"`
}

// OrganizationActivity captures org events.
type OrganizationActivity struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	ActorID        *uint     `json:"actor_id"`
	Action         string    `json:"action" gorm:"not null"`
	Details        string    `json:"details"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// BillingSubscription represents a plan and billing status.
type BillingSubscription struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	Plan           string    `json:"plan" gorm:"default:Starter"`
	Status         string    `json:"status" gorm:"default:ACTIVE"`
	RenewsAt       *time.Time `json:"renews_at"`
	Seats          int       `json:"seats" gorm:"default:1"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BillingInvoice represents a billing invoice.
type BillingInvoice struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	AmountCents    int64     `json:"amount_cents"`
	Currency       string    `json:"currency" gorm:"default:USD"`
	Status         string    `json:"status" gorm:"default:OPEN"`
	PeriodStart    *time.Time `json:"period_start"`
	PeriodEnd      *time.Time `json:"period_end"`
	InvoiceURL     string    `json:"invoice_url"`
	IssuedAt       time.Time `json:"issued_at" gorm:"autoCreateTime"`
}

// BillingPaymentMethod represents a stored payment method.
type BillingPaymentMethod struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	Type           string    `json:"type" gorm:"default:CARD"`
	Last4          string    `json:"last4"`
	Brand          string    `json:"brand"`
	ExpiryMonth    int       `json:"expiry_month"`
	ExpiryYear     int       `json:"expiry_year"`
	IsDefault      bool      `json:"is_default" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// BillingEvent captures billing changes.
type BillingEvent struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	ActorID        *uint     `json:"actor_id"`
	Type           string    `json:"type" gorm:"not null"`
	Details        string    `json:"details"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}
