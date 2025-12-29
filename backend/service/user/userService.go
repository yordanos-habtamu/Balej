package user

import (
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/repository"
)

type UserService interface {
	Create(user *models.User) error
	GetAll() ([]*models.User, error)
   FindByID(id uint) (*models.User, error)
	FindByEmail(email string)(*models.User,error)
	Login(username string,password string) (token string,err error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetAll() ([]*models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) FindByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func(s *userService)Login(email string,password string)(token string ,err error){
	return s.repo.Login(email,password)
}

func(s *userService) FindByEmail(email string) (*models.User,error){
	return s.repo.FindByEmail(email)
}