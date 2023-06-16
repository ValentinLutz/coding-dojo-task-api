package main

import (
	"appchi/incoming/openapi"
	"appchi/incoming/taskapi"
	"appchi/internal/port"
	"appchi/internal/service"
	"appchi/outgoing/taskrepo"
	"errors"
	"log"
	"strconv"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	serverPort, ok := os.LookupEnv("PORT")
	if !ok {
		serverPort = "8080"
	}
	serverPortInt, err := strconv.Atoi(serverPort)
	if err != nil {
		log.Fatalf("failed to parse env PORT: %v", err)
	}

	useInMemory, ok := os.LookupEnv("USE_IN_MEMORY")
	if !ok {
		useInMemory = "true"
	}
	useInMemoryBool, err := strconv.ParseBool(useInMemory)
	if err != nil {
		log.Fatalf("failed to parse env USE_IN_MEMORY: %v", err)
	}

	var taskRepository port.TaskRepository
	if useInMemoryBool {
		taskRepository = taskrepo.NewMemory()
	} else {
		database := NewDatabase()
		taskRepository = taskrepo.NewPostres(database)
	}

	taskService := service.NewTask(taskRepository)
	taskApi := taskapi.New(taskService)

	handler := NewHandler(taskApi)
	server := NewServer(serverPortInt, handler)

	go server.Start()

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received signal: %v\n", (<-stopChannel).String())

	server.Stop()
}

func NewDatabase() *sqlx.DB {
	databaseHost := os.Getenv("POSTGRES_HOST")
	databasePort := os.Getenv("POSTGRES_PORT")
	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")

	psqlInfo := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		databaseHost, databasePort, databaseUser, databasePassword, databaseName,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	initTasksTable(db)

	return db
}

func initTasksTable(db *sqlx.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS public.tasks")
	if err != nil {
		log.Fatalf("failed to drop table public.tasks: %v\n", err)
	}
	log.Println("dropped table public.tasks")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS public.tasks
		( 
			task_id UUID PRIMARY KEY NOT NULL, 
			title TEXT, 
			description TEXT 
		)`,
	)
	if err != nil {
		log.Fatalf("failed to create table public.tasks: %v\n", err)
	}
	log.Println("created table public.tasks")
}

func NewHandler(taskApi http.Handler) http.Handler {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		openAPI := openapi.New()
		openAPI.RegisterRoutes(r)
	})

	router.Group(func(r chi.Router) {
		r.Mount("/", taskApi)
	})

	return router
}

type Server struct {
	httpServer *http.Server
	port       int
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		},
		port: port,
	}
}

func (server *Server) Start() {
	log.Printf("starting server on port: %v\n", server.port)

	err := server.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start server: %s\n", err)
	}
}

func (server *Server) Stop() {
	timeout := time.Second * 20
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("stopping server with timeout: %v\n", timeout.Seconds())
	err := server.httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed to stop server gracefully: %s\n", err)
	}
	log.Println("server stopped")
	// stop other connections like database, message queue
}
