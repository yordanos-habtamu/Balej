package organization

import (
	"fmt"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// OrganizationService manages org profile, members, invites, and billing.
type OrganizationService interface {
	UpsertOrganization(ownerID uint, updates *models.Organization) (*models.Organization, error)
	GetByOwner(ownerID uint) (*models.Organization, error)
	ListMembers(orgID uint) ([]*models.OrganizationMember, error)
	InviteMember(orgID uint, inviterID uint, inviteeID *uint, inviteeEmail string, message string) (*models.OrganizationInvite, error)
	ListInvites(orgID uint) ([]*models.OrganizationInvite, error)
	RespondToInvite(inviteID uint, accept bool) (*models.OrganizationInvite, error)
	RemoveMember(orgID uint, memberID uint) (bool, error)
	UpdateMemberRole(orgID uint, memberID uint, role string) (*models.OrganizationMember, error)
	ListActivity(orgID uint, limit int) ([]*models.OrganizationActivity, error)
	AddActivity(orgID uint, actorID *uint, action string, details string) error
	UpsertSubscription(orgID uint, updates *models.BillingSubscription) (*models.BillingSubscription, error)
	GetSubscription(orgID uint) (*models.BillingSubscription, error)
	ListInvoices(orgID uint) ([]*models.BillingInvoice, error)
	ListPaymentMethods(orgID uint) ([]*models.BillingPaymentMethod, error)
	AddPaymentMethod(orgID uint, method *models.BillingPaymentMethod) (*models.BillingPaymentMethod, error)
	SetDefaultPaymentMethod(orgID uint, methodID uint) (*models.BillingPaymentMethod, error)
	ListBillingEvents(orgID uint) ([]*models.BillingEvent, error)
	CreateBillingEvent(event *models.BillingEvent) (*models.BillingEvent, error)
	DeletePaymentMethod(orgID uint, methodID uint) (bool, error)
	CreateMockInvoice(orgID uint, amountCents int64, currency string) (*models.BillingInvoice, error)
	UpdateInvoiceStatus(orgID uint, invoiceID uint, status string) (*models.BillingInvoice, error)
}

type organizationService struct {
	db *gorm.DB
}

func NewOrganizationService(db *gorm.DB) OrganizationService {
	return &organizationService{db}
}

func (s *organizationService) UpsertOrganization(ownerID uint, updates *models.Organization) (*models.Organization, error) {
	var org models.Organization
	if err := s.db.Where("owner_id = ?", ownerID).First(&org).Error; err != nil {
		org = models.Organization{
			OwnerID:     ownerID,
			Name:        updates.Name,
			Industry:    updates.Industry,
			Size:        updates.Size,
			Website:     updates.Website,
			Description: updates.Description,
		}
		if err := s.db.Create(&org).Error; err != nil {
			return nil, err
		}
		member := &models.OrganizationMember{
			OrganizationID: org.ID,
			UserID:         ownerID,
			Role:           "OWNER",
			JoinedAt:       time.Now(),
		}
		_ = s.db.Create(member).Error
		_ = s.AddActivity(org.ID, &ownerID, "ORG_CREATED", "Organization created")
		return &org, nil
	}
	if updates.Name != "" {
		org.Name = updates.Name
	}
	if updates.Industry != "" {
		org.Industry = updates.Industry
	}
	if updates.Size != "" {
		org.Size = updates.Size
	}
	if updates.Website != "" {
		org.Website = updates.Website
	}
	if updates.Description != "" {
		org.Description = updates.Description
	}
	if err := s.db.Save(&org).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(org.ID, &ownerID, "ORG_UPDATED", "Organization updated")
	return &org, nil
}

func (s *organizationService) GetByOwner(ownerID uint) (*models.Organization, error) {
	var org models.Organization
	if err := s.db.Where("owner_id = ?", ownerID).First(&org).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *organizationService) ListMembers(orgID uint) ([]*models.OrganizationMember, error) {
	var members []*models.OrganizationMember
	if err := s.db.Where("organization_id = ?", orgID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (s *organizationService) InviteMember(orgID uint, inviterID uint, inviteeID *uint, inviteeEmail string, message string) (*models.OrganizationInvite, error) {
	invite := &models.OrganizationInvite{
		OrganizationID: orgID,
		InviterID:      inviterID,
		InviteeID:      inviteeID,
		InviteeEmail:   inviteeEmail,
		Status:         "PENDING",
		Message:        message,
	}
	if err := s.db.Create(invite).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(orgID, &inviterID, "ORG_INVITE_SENT", "Organization invite sent")
	return invite, nil
}

func (s *organizationService) ListInvites(orgID uint) ([]*models.OrganizationInvite, error) {
	var invites []*models.OrganizationInvite
	if err := s.db.Where("organization_id = ?", orgID).Order("created_at desc").Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

func (s *organizationService) RespondToInvite(inviteID uint, accept bool) (*models.OrganizationInvite, error) {
	var invite models.OrganizationInvite
	if err := s.db.First(&invite, inviteID).Error; err != nil {
		return nil, err
	}
	status := "DECLINED"
	if accept {
		status = "ACCEPTED"
	}
	invite.Status = status
	now := time.Now()
	invite.RespondedAt = &now
	if err := s.db.Save(&invite).Error; err != nil {
		return nil, err
	}
	if accept && invite.InviteeID != nil {
		member := &models.OrganizationMember{
			OrganizationID: invite.OrganizationID,
			UserID:         *invite.InviteeID,
			Role:           "MEMBER",
			JoinedAt:       now,
		}
		_ = s.db.Create(member).Error
	}
	action := "ORG_INVITE_DECLINED"
	if accept {
		action = "ORG_INVITE_ACCEPTED"
	}
	_ = s.AddActivity(invite.OrganizationID, invite.InviteeID, action, "Organization invite response")
	return &invite, nil
}

func (s *organizationService) RemoveMember(orgID uint, memberID uint) (bool, error) {
	if err := s.db.Where("organization_id = ?", orgID).Delete(&models.OrganizationMember{}, memberID).Error; err != nil {
		return false, err
	}
	_ = s.AddActivity(orgID, nil, "ORG_MEMBER_REMOVED", "Organization member removed")
	return true, nil
}

func (s *organizationService) UpdateMemberRole(orgID uint, memberID uint, role string) (*models.OrganizationMember, error) {
	var member models.OrganizationMember
	if err := s.db.Where("organization_id = ?", orgID).First(&member, memberID).Error; err != nil {
		return nil, err
	}
	if role != "" {
		member.Role = role
	}
	if err := s.db.Save(&member).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(orgID, nil, "ORG_ROLE_UPDATED", "Organization member role updated")
	return &member, nil
}

func (s *organizationService) ListActivity(orgID uint, limit int) ([]*models.OrganizationActivity, error) {
	if limit <= 0 {
		limit = 50
	}
	var activities []*models.OrganizationActivity
	if err := s.db.Where("organization_id = ?", orgID).Order("created_at desc").Limit(limit).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (s *organizationService) AddActivity(orgID uint, actorID *uint, action string, details string) error {
	entry := &models.OrganizationActivity{
		OrganizationID: orgID,
		ActorID:        actorID,
		Action:         action,
		Details:        details,
	}
	return s.db.Create(entry).Error
}

func (s *organizationService) UpsertSubscription(orgID uint, updates *models.BillingSubscription) (*models.BillingSubscription, error) {
	var sub models.BillingSubscription
	if err := s.db.Where("organization_id = ?", orgID).First(&sub).Error; err != nil {
		sub = models.BillingSubscription{
			OrganizationID: orgID,
			Plan:           updates.Plan,
			Status:         updates.Status,
			RenewsAt:       updates.RenewsAt,
			Seats:          updates.Seats,
		}
		if err := s.db.Create(&sub).Error; err != nil {
			return nil, err
		}
		_ = s.AddActivity(orgID, nil, "BILLING_CREATED", "Billing subscription created")
		_, _ = s.CreateBillingEvent(&models.BillingEvent{
			OrganizationID: orgID,
			Type:           "SUBSCRIPTION_CREATED",
			Details:        sub.Plan,
		})
		_, _ = s.CreateMockInvoice(orgID, 0, "USD")
		return &sub, nil
	}
	if updates.Plan != "" {
		sub.Plan = updates.Plan
	}
	if updates.Status != "" {
		sub.Status = updates.Status
	}
	if updates.RenewsAt != nil {
		sub.RenewsAt = updates.RenewsAt
	}
	if updates.Seats != 0 {
		sub.Seats = updates.Seats
	}
	if err := s.db.Save(&sub).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(orgID, nil, "BILLING_UPDATED", "Billing subscription updated")
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "SUBSCRIPTION_UPDATED",
		Details:        sub.Plan,
	})
	_, _ = s.CreateMockInvoice(orgID, 0, "USD")
	return &sub, nil
}

func (s *organizationService) GetSubscription(orgID uint) (*models.BillingSubscription, error) {
	var sub models.BillingSubscription
	if err := s.db.Where("organization_id = ?", orgID).First(&sub).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

func (s *organizationService) ListInvoices(orgID uint) ([]*models.BillingInvoice, error) {
	var invoices []*models.BillingInvoice
	if err := s.db.Where("organization_id = ?", orgID).Order("issued_at desc").Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (s *organizationService) ListPaymentMethods(orgID uint) ([]*models.BillingPaymentMethod, error) {
	var methods []*models.BillingPaymentMethod
	if err := s.db.Where("organization_id = ?", orgID).Order("created_at desc").Find(&methods).Error; err != nil {
		return nil, err
	}
	return methods, nil
}

func (s *organizationService) AddPaymentMethod(orgID uint, method *models.BillingPaymentMethod) (*models.BillingPaymentMethod, error) {
	method.OrganizationID = orgID
	if method.IsDefault {
		s.db.Model(&models.BillingPaymentMethod{}).Where("organization_id = ?", orgID).Update("is_default", false)
	}
	if err := s.db.Create(method).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(orgID, nil, "PAYMENT_METHOD_ADDED", "Payment method added")
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "PAYMENT_METHOD_ADDED",
		Details:        method.Type,
	})
	return method, nil
}

func (s *organizationService) SetDefaultPaymentMethod(orgID uint, methodID uint) (*models.BillingPaymentMethod, error) {
	var method models.BillingPaymentMethod
	if err := s.db.Where("organization_id = ?", orgID).First(&method, methodID).Error; err != nil {
		return nil, err
	}
	s.db.Model(&models.BillingPaymentMethod{}).Where("organization_id = ?", orgID).Update("is_default", false)
	method.IsDefault = true
	if err := s.db.Save(&method).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(orgID, nil, "PAYMENT_METHOD_DEFAULT", "Default payment method updated")
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "PAYMENT_METHOD_DEFAULT",
		Details:        method.Brand,
	})
	return &method, nil
}

func (s *organizationService) ListBillingEvents(orgID uint) ([]*models.BillingEvent, error) {
	var events []*models.BillingEvent
	if err := s.db.Where("organization_id = ?", orgID).Order("created_at desc").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (s *organizationService) CreateBillingEvent(event *models.BillingEvent) (*models.BillingEvent, error) {
	if err := s.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (s *organizationService) DeletePaymentMethod(orgID uint, methodID uint) (bool, error) {
	if err := s.db.Where("organization_id = ?", orgID).Delete(&models.BillingPaymentMethod{}, methodID).Error; err != nil {
		return false, err
	}
	_ = s.AddActivity(orgID, nil, "PAYMENT_METHOD_REMOVED", "Payment method removed")
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "PAYMENT_METHOD_REMOVED",
		Details:        fmt.Sprintf("MethodID:%d", methodID),
	})
	return true, nil
}

func (s *organizationService) CreateMockInvoice(orgID uint, amountCents int64, currency string) (*models.BillingInvoice, error) {
	if currency == "" {
		currency = "USD"
	}
	if amountCents <= 0 {
		amountCents = s.estimateInvoiceAmount(orgID)
	}
	invoice := &models.BillingInvoice{
		OrganizationID: orgID,
		AmountCents:    amountCents,
		Currency:       currency,
		Status:         "OPEN",
		IssuedAt:       time.Now(),
	}
	if err := s.db.Create(invoice).Error; err != nil {
		return nil, err
	}
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "INVOICE_CREATED",
		Details:        fmt.Sprintf("InvoiceID:%d Amount:%d", invoice.ID, invoice.AmountCents),
	})
	return invoice, nil
}

func (s *organizationService) UpdateInvoiceStatus(orgID uint, invoiceID uint, status string) (*models.BillingInvoice, error) {
	var invoice models.BillingInvoice
	if err := s.db.Where("organization_id = ?", orgID).First(&invoice, invoiceID).Error; err != nil {
		return nil, err
	}
	if !isValidInvoiceTransition(invoice.Status, status) {
		return nil, fmt.Errorf("invalid invoice status transition")
	}
	if status != "" {
		invoice.Status = status
	}
	if err := s.db.Save(&invoice).Error; err != nil {
		return nil, err
	}
	_, _ = s.CreateBillingEvent(&models.BillingEvent{
		OrganizationID: orgID,
		Type:           "INVOICE_STATUS_UPDATED",
		Details:        fmt.Sprintf("InvoiceID:%d Status:%s", invoiceID, status),
	})
	return &invoice, nil
}

func (s *organizationService) estimateInvoiceAmount(orgID uint) int64 {
	// Simple mock pricing: $29 per seat
	var sub models.BillingSubscription
	if err := s.db.Where("organization_id = ?", orgID).First(&sub).Error; err == nil {
		seats := sub.Seats
		if seats <= 0 {
			seats = 1
		}
		return int64(seats) * 2900
	}
	return 2900
}

func isValidInvoiceTransition(current string, next string) bool {
	if next == "" || current == next {
		return true
	}
	switch current {
	case "OPEN":
		return next == "PAID" || next == "VOID"
	case "PAID":
		return false
	case "VOID":
		return false
	default:
		return next == "OPEN" || next == "PAID" || next == "VOID"
	}
}
