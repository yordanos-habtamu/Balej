package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	pubsub "github.com/yordanos-habtamu/GormWithGraphql/pubsub"
	"github.com/yordanos-habtamu/GormWithGraphql/service/job"
	"github.com/yordanos-habtamu/GormWithGraphql/service/user"
)

type Resolver struct {
	UserService user.UserService
	JobService  job.JobService
	ChatService user.ChatService
	PubSub      *pubsub.PubSub
}
