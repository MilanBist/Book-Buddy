package api

import (
	"net/http"
	"time"

	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/middlewares"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	askllm "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/askLLm"
	extractanswer "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractAnswer"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractpdf"
	storeextracteddata "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/storeExtractedData"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qdrant/go-client/qdrant"
)

type Handler struct{
	// declare the models here
	server *models.Server
}

func NewServer(vectorStore *qdrant.Client, cfg *config.Config, postgres *pgxpool.Pool) *models.Server{
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 *time.Second))

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
  }))

  	server := &models.Server{
		Router: r,
		Store: vectorStore,
		Config: cfg,
		PostgresDB: postgres,
	}

	handler := &Handler{
		server: server,
	}

	handler.setupRoutes()
	return handler.server
}

func(h *Handler) setupRoutes(){
	// initally test with ping pong statement
	// create handlers for each of the endpoints
	store := &dbqueries.PostgresStore{
		Db: h.server.PostgresDB,
	}
	token := &utils.Token{
		AccessToken: "",
	}
	embed := &extractanswer.Embeddings{
		Server: h.server,
	}
	llm := &askllm.LLMWork{
		Server: h.server,
	}
	storeFile := &storeextracteddata.StoreDataToLocation{
		BasePath: "./uploadedFiles",
	}
	dataStore := &extractpdf.DataExtract{
		Server: h.server,
	}

	loginHandler := &LoginHandler{
		Check: store,
		Token: token,
	}
	registerHandler := &RegisterHandler{
		Register: store,
		Token: token,
	}
	questionHandler := &QuestionHanlder{
		StoreQuestion: store,
		AnswerQuestion: embed,
		Response: llm,
	}
	bookHandler := &BookGettingHandler{
		Book: store,
	}
	conversationHandler := &UserConversationHandler{
		Conversation: store,
	}

	bookUploadHandler := &RawPdfHandler{
		StoreBook: store,
		GetFileLocation: storeFile,
		ExtractPdf: dataStore,
	}




	h.server.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("pong"));
	})
	h.server.Router.Route(
		"/api", func(r chi.Router) {
			r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Success in API call!"))
			})

			// for the login and registration
			r.Post("/login", loginHandler.HandleLogin)
			r.Post("/register", registerHandler.HandleRegister)

			// create the protected handlers
			r.Group(func(r chi.Router){
				// add the middlewares here
				r.Use(middlewares.LoggingMiddleware)
				r.Use(middlewares.AuthMiddleware)

				r.Post("/handlePdf", bookUploadHandler.HandleRawPdfFile)
				r.Post("/extractAnswer", questionHandler.HandleRawQuestion)

				r.Get("/getBooks", bookHandler.HandleGettingBooks)
				r.Get("/getConversation", conversationHandler.HandleConversation)
			})

			
		})
}