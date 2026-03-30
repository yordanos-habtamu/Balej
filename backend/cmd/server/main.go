package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	// Your local imports, adjust paths as needed
	"github.com/yordanos-habtamu/GormWithGraphql/graph"
	"github.com/yordanos-habtamu/GormWithGraphql/middleware"
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/models/db"
	pubsub "github.com/yordanos-habtamu/GormWithGraphql/pubsub"
	"github.com/yordanos-habtamu/GormWithGraphql/repository"
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

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to database
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	database.AutoMigrate(&models.User{})

	// Dependency injection setup (your services)
	userRepo := repository.NewUserRepository(database)
	userService := user.NewUserService(userRepo)
	adminService := admin.NewAdminService(database)

	jobRepo := repository.NewjobRepository(database)
	jobService := job.NewJobService(jobRepo)

	chatRepo := repository.NewChatRepository(database)
	chatService := user.NewChatService(chatRepo)

	notificationService := notification.NewNotificationService(database)
	moderationService := moderation.NewModerationService(database)
	secureRelayService := security.NewSecureRelayService()
	complianceService := compliance.NewComplianceService(database)
	communicationService := communication.NewCommunicationService(database)
	organizationService := organization.NewOrganizationService(database)
	rbacService := rbac.NewRBACService(database)
	profileService := profile.NewProfileService(database)
	projectService := project.NewProjectService(database)

	pubsub := pubsub.NewPubSub()

	resolver := &graph.Resolver{
		UserService: userService,
		AdminService: adminService,
		JobService:  jobService,
		ChatService: chatService,
		NotificationService: notificationService,
		ModerationService: moderationService,
		SecureRelayService: secureRelayService,
		ComplianceService: complianceService,
		CommunicationService: communicationService,
		OrganizationService: organizationService,
		RBACService: rbacService,
		ProfileService: profileService,
		ProjectService: projectService,
		PubSub:      pubsub,
	}

	// Start mock compliance job worker
	stopComplianceWorker := make(chan struct{})
	go complianceService.StartAutoCompleteWorker(stopComplianceWorker, 30*time.Second)

	// Setup gqlgen server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Add WebSocket transport for subscriptions
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Change this in production to restrict origins
				return true
			},
		},
	})

	// Add POST support for queries/mutations
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	// Setup router
	r := mux.NewRouter()
	r.Handle("/query", middleware.AuthMiddleware(srv)).Methods("GET", "POST", "OPTIONS")
	r.Handle("/", playground.Handler("GraphQL Playground", "/query"))

	// CORS middleware config — replace with your frontend URL
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:8000",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Upgrade",
			"Connection",
			"Sec-Websocket-Key",
			"Sec-Websocket-Version",
			"Sec-Websocket-Extensions",
		}),
		handlers.AllowCredentials(),
	)(r)

	log.Printf("🚀 Server ready at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
