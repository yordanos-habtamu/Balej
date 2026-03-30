package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	pubsub "github.com/yordanos-habtamu/GormWithGraphql/pubsub"
	"github.com/yordanos-habtamu/GormWithGraphql/service/admin"
	"github.com/yordanos-habtamu/GormWithGraphql/service/compliance"
	"github.com/yordanos-habtamu/GormWithGraphql/service/communication"
	"github.com/yordanos-habtamu/GormWithGraphql/service/job"
	"github.com/yordanos-habtamu/GormWithGraphql/service/moderation"
	"github.com/yordanos-habtamu/GormWithGraphql/service/notification"
	"github.com/yordanos-habtamu/GormWithGraphql/service/organization"
	"github.com/yordanos-habtamu/GormWithGraphql/service/project"
	"github.com/yordanos-habtamu/GormWithGraphql/service/profile"
	"github.com/yordanos-habtamu/GormWithGraphql/service/rbac"
	"github.com/yordanos-habtamu/GormWithGraphql/service/security"
	"github.com/yordanos-habtamu/GormWithGraphql/service/user"
)

type Resolver struct {
	UserService         user.UserService
	AdminService        admin.AdminService
	JobService          job.JobService
	ChatService         user.ChatService
	NotificationService notification.NotificationService
	ModerationService   moderation.ModerationService
	SecureRelayService  security.SecureRelayService
	ComplianceService   compliance.ComplianceService
	CommunicationService communication.CommunicationService
	OrganizationService organization.OrganizationService
	RBACService         rbac.RBACService
	ProfileService      profile.ProfileService
	ProjectService      project.ProjectService
	PubSub              *pubsub.PubSub
}
