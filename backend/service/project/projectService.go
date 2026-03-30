package project

import (
	"fmt"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// ProjectService manages projects and invitations.
type ProjectService interface {
	CreateProject(ownerID uint, input *models.Project) (*models.Project, error)
	UpdateProject(projectID uint, updates *models.Project) (*models.Project, error)
	DeleteProject(projectID uint) (bool, error)
	GetProject(projectID uint) (*models.Project, error)
	ListProjectsForUser(userID uint) ([]*models.Project, error)
	InviteToProject(projectID uint, inviterID uint, inviteeID *uint, inviteeEmail string, role string, message string) (*models.ProjectInvitation, error)
	GetInvitation(invitationID uint) (*models.ProjectInvitation, error)
	ListInvitationsForUser(userID uint) ([]*models.ProjectInvitation, error)
	ListInvitationsForProject(projectID uint) ([]*models.ProjectInvitation, error)
	RespondToInvitation(invitationID uint, accept bool) (*models.ProjectInvitation, error)
	RemoveProjectMember(projectID uint, memberID uint) (bool, error)
	IsProjectMember(projectID uint, userID uint) (bool, error)
	GetMemberRole(projectID uint, userID uint) (string, error)
	ResendInvitation(invitationID uint) (*models.ProjectInvitation, error)
	UpdateMemberRole(projectID uint, memberID uint, role string) (*models.ProjectMember, error)
	ListProjectActivities(projectID uint, limit int) ([]*models.ProjectActivity, error)
	AddActivity(projectID uint, actorID *uint, action string, details string) error
	CreateTask(task *models.ProjectTask) (*models.ProjectTask, error)
	GetTask(taskID uint) (*models.ProjectTask, error)
	UpdateTask(taskID uint, updates *models.ProjectTask) (*models.ProjectTask, error)
	DeleteTask(taskID uint) (bool, error)
	ListTasks(projectID uint) ([]*models.ProjectTask, error)
	CreateDeliverable(deliverable *models.ProjectDeliverable) (*models.ProjectDeliverable, error)
	GetDeliverable(deliverableID uint) (*models.ProjectDeliverable, error)
	UpdateDeliverable(deliverableID uint, updates *models.ProjectDeliverable) (*models.ProjectDeliverable, error)
	DeleteDeliverable(deliverableID uint) (bool, error)
	ListDeliverables(projectID uint) ([]*models.ProjectDeliverable, error)
	CreateSubmission(submission *models.ProjectSubmission) (*models.ProjectSubmission, error)
	GetSubmission(submissionID uint) (*models.ProjectSubmission, error)
	ReviewSubmission(submissionID uint, status string, feedback string, reviewerID uint) (*models.ProjectSubmission, error)
	ListSubmissions(projectID uint, deliverableID *uint) ([]*models.ProjectSubmission, error)
	CreateProjectChat(projectID uint, chatID string) (*models.ProjectChat, error)
	ListProjectChats(projectID uint) ([]*models.ProjectChat, error)
	IncrementProjectView(projectID uint) error
}

type projectService struct {
	db *gorm.DB
}

// NewProjectService initializes project service.
func NewProjectService(db *gorm.DB) ProjectService {
	return &projectService{db}
}

// CreateProject creates a project and adds the owner as a member.
func (s *projectService) CreateProject(ownerID uint, input *models.Project) (*models.Project, error) {
	project := &models.Project{
		OwnerID:     ownerID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		IsPublic:    input.IsPublic,
		CoverImage:  input.CoverImage,
		Summary:     input.Summary,
	}
	if project.Status == "" {
		project.Status = "ACTIVE"
	}
	if err := s.db.Create(project).Error; err != nil {
		return nil, err
	}
	member := &models.ProjectMember{
		ProjectID: project.ID,
		UserID:    ownerID,
		Role:      "OWNER",
		JoinedAt:  time.Now(),
	}
	if err := s.db.Create(member).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(project.ID, &ownerID, "PROJECT_CREATED", "Project created")
	return s.GetProject(project.ID)
}

// UpdateProject updates project metadata.
func (s *projectService) UpdateProject(projectID uint, updates *models.Project) (*models.Project, error) {
	var project models.Project
	if err := s.db.First(&project, projectID).Error; err != nil {
		return nil, err
	}
	if updates.Title != "" {
		project.Title = updates.Title
	}
	if updates.Description != "" {
		project.Description = updates.Description
	}
	if updates.Status != "" {
		project.Status = updates.Status
	}
	project.IsPublic = updates.IsPublic
	if updates.CoverImage != "" {
		project.CoverImage = updates.CoverImage
	}
	if updates.Summary != "" {
		project.Summary = updates.Summary
	}
	if err := s.db.Save(&project).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(project.ID, &project.OwnerID, "PROJECT_UPDATED", "Project updated")
	return s.GetProject(project.ID)
}

// DeleteProject removes a project and its related data.
func (s *projectService) DeleteProject(projectID uint) (bool, error) {
	if err := s.db.Delete(&models.Project{}, projectID).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetProject returns a project with relationships loaded.
func (s *projectService) GetProject(projectID uint) (*models.Project, error) {
	var project models.Project
	err := s.db.Preload("Owner").
		Preload("Members").
		Preload("Members.User").
		Preload("Invitations").
		First(&project, projectID).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// ListProjectsForUser returns projects owned by or shared with a user.
func (s *projectService) ListProjectsForUser(userID uint) ([]*models.Project, error) {
	var projects []*models.Project
	sub := s.db.Table("project_members").Select("project_id").Where("user_id = ?", userID)
	if err := s.db.Preload("Owner").
		Preload("Members").
		Preload("Members.User").
		Preload("Invitations").
		Where("owner_id = ? OR id IN (?)", userID, sub).
		Order("updated_at desc").
		Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// InviteToProject creates an invitation record.
func (s *projectService) InviteToProject(projectID uint, inviterID uint, inviteeID *uint, inviteeEmail string, role string, message string) (*models.ProjectInvitation, error) {
	var existing models.ProjectInvitation
	query := s.db.Where("project_id = ? AND status = ?", projectID, "PENDING")
	if inviteeID != nil {
		query = query.Where("invitee_id = ?", *inviteeID)
	} else if inviteeEmail != "" {
		query = query.Where("invitee_email = ?", inviteeEmail)
	}
	if err := query.First(&existing).Error; err == nil {
		return nil, fmt.Errorf("invitation already pending")
	}

	invite := &models.ProjectInvitation{
		ProjectID:    projectID,
		InviterID:    inviterID,
		InviteeID:    inviteeID,
		InviteeEmail: inviteeEmail,
		Role:         role,
		Status:       "PENDING",
		Message:      message,
	}
	if invite.Role == "" {
		invite.Role = "COLLABORATOR"
	}
	expires := time.Now().Add(7 * 24 * time.Hour)
	invite.ExpiresAt = &expires
	if err := s.db.Create(invite).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(projectID, &inviterID, "INVITE_SENT", "Project invitation sent")
	return invite, nil
}

// GetInvitation fetches a single invitation.
func (s *projectService) GetInvitation(invitationID uint) (*models.ProjectInvitation, error) {
	var invite models.ProjectInvitation
	if err := s.db.First(&invite, invitationID).Error; err != nil {
		return nil, err
	}
	return &invite, nil
}

// ListInvitationsForUser lists invitations for a user.
func (s *projectService) ListInvitationsForUser(userID uint) ([]*models.ProjectInvitation, error) {
	var invites []*models.ProjectInvitation
	if err := s.db.Where("invitee_id = ?", userID).Order("created_at desc").Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

// ListInvitationsForProject returns invitations for a project.
func (s *projectService) ListInvitationsForProject(projectID uint) ([]*models.ProjectInvitation, error) {
	var invites []*models.ProjectInvitation
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

// RespondToInvitation accepts or declines an invitation.
func (s *projectService) RespondToInvitation(invitationID uint, accept bool) (*models.ProjectInvitation, error) {
	var invite models.ProjectInvitation
	if err := s.db.First(&invite, invitationID).Error; err != nil {
		return nil, err
	}
	if invite.ExpiresAt != nil && invite.ExpiresAt.Before(time.Now()) && invite.Status == "PENDING" {
		invite.Status = "EXPIRED"
		now := time.Now()
		invite.RespondedAt = &now
		_ = s.db.Save(&invite).Error
		return nil, fmt.Errorf("invitation expired")
	}
	status := "DECLINED"
	if accept {
		status = "ACCEPTED"
	}
	invite.Status = status
	now := time.Now()
	invite.RespondedAt = &now

	tx := s.db.Begin()
	if err := tx.Save(&invite).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if accept && invite.InviteeID != nil {
		var existing models.ProjectMember
		if err := tx.Where("project_id = ? AND user_id = ?", invite.ProjectID, *invite.InviteeID).
			First(&existing).Error; err == gorm.ErrRecordNotFound {
			member := &models.ProjectMember{
				ProjectID: invite.ProjectID,
				UserID:    *invite.InviteeID,
				Role:      invite.Role,
				JoinedAt:  now,
			}
			if err := tx.Create(member).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		} else if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	action := "INVITE_DECLINED"
	if accept {
		action = "INVITE_ACCEPTED"
	}
	_ = s.AddActivity(invite.ProjectID, invite.InviteeID, action, "Project invitation response")
	return &invite, nil
}

// RemoveProjectMember removes a member from a project.
func (s *projectService) RemoveProjectMember(projectID uint, memberID uint) (bool, error) {
	if err := s.db.Where("project_id = ?", projectID).Delete(&models.ProjectMember{}, memberID).Error; err != nil {
		return false, err
	}
	_ = s.AddActivity(projectID, nil, "MEMBER_REMOVED", "Project member removed")
	return true, nil
}

// IsProjectMember checks if a user is a member of a project.
func (s *projectService) IsProjectMember(projectID uint, userID uint) (bool, error) {
	var count int64
	if err := s.db.Model(&models.ProjectMember{}).Where("project_id = ? AND user_id = ?", projectID, userID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetMemberRole returns the role for a user in a project.
func (s *projectService) GetMemberRole(projectID uint, userID uint) (string, error) {
	var member models.ProjectMember
	if err := s.db.Where("project_id = ? AND user_id = ?", projectID, userID).
		First(&member).Error; err != nil {
		return "", err
	}
	return member.Role, nil
}

// ResendInvitation refreshes expiry and bumps resend counter.
func (s *projectService) ResendInvitation(invitationID uint) (*models.ProjectInvitation, error) {
	var invite models.ProjectInvitation
	if err := s.db.First(&invite, invitationID).Error; err != nil {
		return nil, err
	}
	invite.ResendCount += 1
	expires := time.Now().Add(7 * 24 * time.Hour)
	invite.ExpiresAt = &expires
	invite.Status = "PENDING"
	if err := s.db.Save(&invite).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(invite.ProjectID, &invite.InviterID, "INVITE_RESENT", "Project invitation resent")
	return &invite, nil
}

// UpdateMemberRole updates a project member role.
func (s *projectService) UpdateMemberRole(projectID uint, memberID uint, role string) (*models.ProjectMember, error) {
	var member models.ProjectMember
	if err := s.db.Where("project_id = ?", projectID).First(&member, memberID).Error; err != nil {
		return nil, err
	}
	if role != "" {
		member.Role = role
	}
	if err := s.db.Save(&member).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(projectID, nil, "ROLE_UPDATED", "Project member role updated")
	return &member, nil
}

// ListProjectActivities returns latest project activities.
func (s *projectService) ListProjectActivities(projectID uint, limit int) ([]*models.ProjectActivity, error) {
	if limit <= 0 {
		limit = 50
	}
	var activities []*models.ProjectActivity
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Limit(limit).
		Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

// AddActivity stores a project activity entry.
func (s *projectService) AddActivity(projectID uint, actorID *uint, action string, details string) error {
	entry := &models.ProjectActivity{
		ProjectID: projectID,
		ActorID:   actorID,
		Action:    action,
		Details:   details,
	}
	return s.db.Create(entry).Error
}

// CreateTask creates a project task.
func (s *projectService) CreateTask(task *models.ProjectTask) (*models.ProjectTask, error) {
	if task.Status == "" {
		task.Status = "OPEN"
	}
	if task.Priority == "" {
		task.Priority = "MEDIUM"
	}
	if err := s.db.Create(task).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(task.ProjectID, nil, "TASK_CREATED", "Task created")
	return task, nil
}

// GetTask returns a task by ID.
func (s *projectService) GetTask(taskID uint) (*models.ProjectTask, error) {
	var task models.ProjectTask
	if err := s.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates a task.
func (s *projectService) UpdateTask(taskID uint, updates *models.ProjectTask) (*models.ProjectTask, error) {
	var task models.ProjectTask
	if err := s.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	if updates.Title != "" {
		task.Title = updates.Title
	}
	if updates.Description != "" {
		task.Description = updates.Description
	}
	if updates.Status != "" {
		task.Status = updates.Status
	}
	if updates.Priority != "" {
		task.Priority = updates.Priority
	}
	if updates.AssigneeID != nil {
		task.AssigneeID = updates.AssigneeID
	}
	if updates.DueDate != nil {
		task.DueDate = updates.DueDate
	}
	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(task.ProjectID, nil, "TASK_UPDATED", "Task updated")
	return &task, nil
}

// DeleteTask removes a task.
func (s *projectService) DeleteTask(taskID uint) (bool, error) {
	var task models.ProjectTask
	if err := s.db.First(&task, taskID).Error; err != nil {
		return false, err
	}
	if err := s.db.Delete(&models.ProjectTask{}, taskID).Error; err != nil {
		return false, err
	}
	_ = s.AddActivity(task.ProjectID, nil, "TASK_DELETED", "Task deleted")
	return true, nil
}

// ListTasks returns tasks for a project.
func (s *projectService) ListTasks(projectID uint) ([]*models.ProjectTask, error) {
	var tasks []*models.ProjectTask
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// CreateDeliverable creates a deliverable.
func (s *projectService) CreateDeliverable(deliverable *models.ProjectDeliverable) (*models.ProjectDeliverable, error) {
	if deliverable.Status == "" {
		deliverable.Status = "PENDING"
	}
	if err := s.db.Create(deliverable).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(deliverable.ProjectID, nil, "DELIVERABLE_CREATED", "Deliverable created")
	return deliverable, nil
}

// GetDeliverable returns a deliverable by ID.
func (s *projectService) GetDeliverable(deliverableID uint) (*models.ProjectDeliverable, error) {
	var deliverable models.ProjectDeliverable
	if err := s.db.First(&deliverable, deliverableID).Error; err != nil {
		return nil, err
	}
	return &deliverable, nil
}

// UpdateDeliverable updates a deliverable.
func (s *projectService) UpdateDeliverable(deliverableID uint, updates *models.ProjectDeliverable) (*models.ProjectDeliverable, error) {
	var deliverable models.ProjectDeliverable
	if err := s.db.First(&deliverable, deliverableID).Error; err != nil {
		return nil, err
	}
	if updates.Title != "" {
		deliverable.Title = updates.Title
	}
	if updates.Description != "" {
		deliverable.Description = updates.Description
	}
	if updates.Status != "" {
		deliverable.Status = updates.Status
	}
	if updates.DueDate != nil {
		deliverable.DueDate = updates.DueDate
	}
	if err := s.db.Save(&deliverable).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(deliverable.ProjectID, nil, "DELIVERABLE_UPDATED", "Deliverable updated")
	return &deliverable, nil
}

// DeleteDeliverable removes a deliverable.
func (s *projectService) DeleteDeliverable(deliverableID uint) (bool, error) {
	var deliverable models.ProjectDeliverable
	if err := s.db.First(&deliverable, deliverableID).Error; err != nil {
		return false, err
	}
	if err := s.db.Delete(&models.ProjectDeliverable{}, deliverableID).Error; err != nil {
		return false, err
	}
	_ = s.AddActivity(deliverable.ProjectID, nil, "DELIVERABLE_DELETED", "Deliverable deleted")
	return true, nil
}

// ListDeliverables returns deliverables for a project.
func (s *projectService) ListDeliverables(projectID uint) ([]*models.ProjectDeliverable, error) {
	var deliverables []*models.ProjectDeliverable
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Find(&deliverables).Error; err != nil {
		return nil, err
	}
	return deliverables, nil
}

// CreateSubmission creates a submission.
func (s *projectService) CreateSubmission(submission *models.ProjectSubmission) (*models.ProjectSubmission, error) {
	if submission.Status == "" {
		submission.Status = "SUBMITTED"
	}
	if err := s.db.Create(submission).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(submission.ProjectID, &submission.SubmitterID, "SUBMISSION_CREATED", "Deliverable submitted")
	return submission, nil
}

// GetSubmission returns a submission by ID.
func (s *projectService) GetSubmission(submissionID uint) (*models.ProjectSubmission, error) {
	var submission models.ProjectSubmission
	if err := s.db.First(&submission, submissionID).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

// ReviewSubmission reviews a submission.
func (s *projectService) ReviewSubmission(submissionID uint, status string, feedback string, reviewerID uint) (*models.ProjectSubmission, error) {
	var submission models.ProjectSubmission
	if err := s.db.First(&submission, submissionID).Error; err != nil {
		return nil, err
	}
	if status != "" {
		submission.Status = status
	}
	submission.Feedback = feedback
	submission.ReviewedBy = &reviewerID
	now := time.Now()
	submission.ReviewedAt = &now
	if err := s.db.Save(&submission).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(submission.ProjectID, &reviewerID, "SUBMISSION_REVIEWED", status)
	return &submission, nil
}

// ListSubmissions returns submissions for a project.
func (s *projectService) ListSubmissions(projectID uint, deliverableID *uint) ([]*models.ProjectSubmission, error) {
	var submissions []*models.ProjectSubmission
	query := s.db.Where("project_id = ?", projectID)
	if deliverableID != nil {
		query = query.Where("deliverable_id = ?", *deliverableID)
	}
	if err := query.Order("created_at desc").Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}

func (s *projectService) CreateProjectChat(projectID uint, chatID string) (*models.ProjectChat, error) {
	link := &models.ProjectChat{
		ProjectID: projectID,
		ChatID:    chatID,
	}
	if err := s.db.Create(link).Error; err != nil {
		return nil, err
	}
	_ = s.AddActivity(projectID, nil, "PROJECT_CHAT_CREATED", "Project chat created")
	return link, nil
}

func (s *projectService) ListProjectChats(projectID uint) ([]*models.ProjectChat, error) {
	var chats []*models.ProjectChat
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func (s *projectService) IncrementProjectView(projectID uint) error {
	return s.db.Model(&models.Project{}).Where("id = ?", projectID).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}
