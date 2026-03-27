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
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/models/db"
	pubsub "github.com/yordanos-habtamu/GormWithGraphql/pubsub"
	"github.com/yordanos-habtamu/GormWithGraphql/repository"
	"github.com/yordanos-habtamu/GormWithGraphql/service/job"
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

	jobRepo := repository.NewjobRepository(database)
	jobService := job.NewJobService(jobRepo)

	chatRepo := repository.NewChatRepository(database)
	chatService := user.NewChatService(chatRepo)

	pubsub := pubsub.NewPubSub()

	resolver := &graph.Resolver{
		UserService: userService,
		JobService:  jobService,
		ChatService: chatService,
		PubSub:      pubsub,
	}

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
	r.Handle("/query", srv).Methods("GET", "POST", "OPTIONS")
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
