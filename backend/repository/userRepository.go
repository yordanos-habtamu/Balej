package repository

import (
	"errors"
	"fmt"


	"github.com/yordanos-habtamu/GormWithGraphql/auth/helpers"
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]*models.User, error)
	FindByID(id uint) (*models.User, error)
	Login( username string, password string) (token string,err error)
	Register(user *models.User) error
	FindByEmail(email string) (*models.User,error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}
func (r *userRepo) FindByEmail(email string) (*models.User,error){
	var user models.User
	  if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, fmt.Errorf("the username was not found: %v", err)
    }
	return &user,nil
}

func (r *userRepo) Login(email string, password string) (token string, err error) {
	u, err := r.FindByEmail(email)
	if err != nil {
		return "", errors.New("no email was found")
	}
	if !helpers.ComparePassword(u.Password, []byte(password)) {
		return "", errors.New("invalid credentials")
	}
	token, err = helpers.GenerateJWT(uint(u.ID), u.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ( r *userRepo) Register(user *models.User) error{
	return r.db.Create(user).Error
}