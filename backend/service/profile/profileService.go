package profile

import (
	"fmt"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// ProfileService manages talent profile data.
type ProfileService interface {
	GetByUserID(userID uint) (*models.TalentProfile, error)
	UpsertProfile(userID uint, updates *models.TalentProfile) (*models.TalentProfile, error)
	SetPublic(userID uint, isPublic bool) (*models.TalentProfile, error)
	AddSkill(profileID uint, skill *models.Skill) (*models.Skill, error)
	UpdateSkill(id uint, updates *models.Skill) (*models.Skill, error)
	DeleteSkill(id uint) (bool, error)
	AddExperience(profileID uint, exp *models.Experience) (*models.Experience, error)
	UpdateExperience(id uint, updates *models.Experience, isCurrent *bool) (*models.Experience, error)
	DeleteExperience(id uint) (bool, error)
	AddWorkSample(profileID uint, ws *models.WorkSample) (*models.WorkSample, error)
	UpdateWorkSample(id uint, updates *models.WorkSample) (*models.WorkSample, error)
	DeleteWorkSample(id uint) (bool, error)
}

type profileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) ProfileService {
	return &profileService{db}
}

func (s *profileService) GetByUserID(userID uint) (*models.TalentProfile, error) {
	var profile models.TalentProfile
	err := s.db.Preload("Skills").Preload("Experiences").Preload("WorkSamples").
		Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *profileService) UpsertProfile(userID uint, updates *models.TalentProfile) (*models.TalentProfile, error) {
	var profile models.TalentProfile
	err := s.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			profile = models.TalentProfile{
				UserID:            userID,
				ProfessionalTitle: updates.ProfessionalTitle,
				Bio:               updates.Bio,
				Location:          updates.Location,
				ProfileImageURL:   updates.ProfileImageURL,
				CompletionPercent: updates.CompletionPercent,
			}
			if err := s.db.Create(&profile).Error; err != nil {
				return nil, err
			}
			return &profile, nil
		}
		return nil, err
	}
	if updates.ProfessionalTitle != "" {
		profile.ProfessionalTitle = updates.ProfessionalTitle
	}
	if updates.Bio != "" {
		profile.Bio = updates.Bio
	}
	if updates.Location != "" {
		profile.Location = updates.Location
	}
	if updates.ProfileImageURL != "" {
		profile.ProfileImageURL = updates.ProfileImageURL
	}
	if updates.CompletionPercent != 0 {
		profile.CompletionPercent = updates.CompletionPercent
	}
	if updates.IsPublic {
		profile.IsPublic = updates.IsPublic
	}
	if err := s.db.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *profileService) SetPublic(userID uint, isPublic bool) (*models.TalentProfile, error) {
	var profile models.TalentProfile
	if err := s.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	profile.IsPublic = isPublic
	if err := s.db.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *profileService) AddSkill(profileID uint, skill *models.Skill) (*models.Skill, error) {
	skill.ProfileID = profileID
	if err := s.db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (s *profileService) UpdateSkill(id uint, updates *models.Skill) (*models.Skill, error) {
	var skill models.Skill
	if err := s.db.First(&skill, id).Error; err != nil {
		return nil, err
	}
	if updates.SkillName != "" {
		skill.SkillName = updates.SkillName
	}
	if updates.ProficiencyLevel != "" {
		skill.ProficiencyLevel = updates.ProficiencyLevel
	}
	if updates.Category != "" {
		skill.Category = updates.Category
	}
	if err := s.db.Save(&skill).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func (s *profileService) DeleteSkill(id uint) (bool, error) {
	if err := s.db.Delete(&models.Skill{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *profileService) AddExperience(profileID uint, exp *models.Experience) (*models.Experience, error) {
	exp.ProfileID = profileID
	if err := s.db.Create(exp).Error; err != nil {
		return nil, err
	}
	return exp, nil
}

func (s *profileService) UpdateExperience(id uint, updates *models.Experience, isCurrent *bool) (*models.Experience, error) {
	var exp models.Experience
	if err := s.db.First(&exp, id).Error; err != nil {
		return nil, err
	}
	if updates.JobTitle != "" {
		exp.JobTitle = updates.JobTitle
	}
	if updates.CompanyName != "" {
		exp.CompanyName = updates.CompanyName
	}
	if !updates.StartDate.IsZero() {
		exp.StartDate = updates.StartDate
	}
	if updates.EndDate != nil {
		exp.EndDate = updates.EndDate
	}
	if updates.Description != "" {
		exp.Description = updates.Description
	}
	if isCurrent != nil {
		exp.IsCurrent = *isCurrent
	}
	if err := s.db.Save(&exp).Error; err != nil {
		return nil, err
	}
	return &exp, nil
}

func (s *profileService) DeleteExperience(id uint) (bool, error) {
	if err := s.db.Delete(&models.Experience{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *profileService) AddWorkSample(profileID uint, ws *models.WorkSample) (*models.WorkSample, error) {
	ws.ProfileID = profileID
	if err := s.db.Create(ws).Error; err != nil {
		return nil, err
	}
	return ws, nil
}

func (s *profileService) UpdateWorkSample(id uint, updates *models.WorkSample) (*models.WorkSample, error) {
	var ws models.WorkSample
	if err := s.db.First(&ws, id).Error; err != nil {
		return nil, err
	}
	if updates.Title != "" {
		ws.Title = updates.Title
	}
	if updates.Description != "" {
		ws.Description = updates.Description
	}
	if updates.ImageURL != "" {
		ws.ImageURL = updates.ImageURL
	}
	if updates.ExternalLink != "" {
		ws.ExternalLink = updates.ExternalLink
	}
	if updates.DisplayOrder != 0 {
		ws.DisplayOrder = updates.DisplayOrder
	}
	if err := s.db.Save(&ws).Error; err != nil {
		return nil, err
	}
	return &ws, nil
}

func (s *profileService) DeleteWorkSample(id uint) (bool, error) {
	if err := s.db.Delete(&models.WorkSample{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// helper for ownership checks
func EnsureProfileOwner(profile *models.TalentProfile, userID uint) error {
	if profile.UserID != userID {
		return fmt.Errorf("unauthorized: profile access denied")
	}
	return nil
}
