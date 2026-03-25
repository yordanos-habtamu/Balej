package user

import(
	 "github.com/yordanos-habtamu/GormWithGraphql/models"
	 "github.com/yordanos-habtamu/GormWithGraphql/repository"
	)

type ChatService interface{
	Create(senderID uint, receiverID uint) (*models.Chat, error)
	GetByID(chatID string) (*models.Chat, error)
	SendMessage(chatID string, message *models.Message , senderId int, recieverId int) error
	GetUserChats(userID uint) ([]*models.Chat, error)
	GetChatHistory(chatID string) ([]*models.Message, error)
}


type chatService struct {
	repo repository.ChatRepository
}
func NewChatService(r repository.ChatRepository) ChatService {
	return &chatService{r}
}
func (s *chatService) Create(senderID uint, receiverID uint) (*models.Chat, error) {	
 chat,err:=	s.repo.Create(senderID,receiverID)
 if err != nil{
	return nil,err
 }
 return chat,nil
}

func (s *chatService) GetByID(chatID string) (*models.Chat, error) {
	chat,err:= s.repo.GetByID(chatID)
	if err!= nil {
		return nil,err
	}
   return chat,nil
}

func (s *chatService) SendMessage(chatID string, message *models.Message,senderId int, recieverId int) error {
	if err := s.repo.SendMessage(chatID, message,senderId,recieverId); err != nil {
		return err
	}
	return nil
}

func (s *chatService) GetUserChats(userID uint) ([]*models.Chat, error) {
	return s.repo.GetChatsByUserID(userID)
}

func (s *chatService) GetChatHistory(chatID string) ([]*models.Message, error) {
	return s.repo.GetMessagesByChatID(chatID)
}