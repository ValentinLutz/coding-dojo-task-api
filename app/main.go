package main

import (
	"app/incoming/openapi"
	"app/incoming/taskapi"
	"app/internal/task"
	internalTask "app/internal/task"
	"errors"
	"log"

	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	port   = *flag.Int("port", 8080, "http server port")
	memory = *flag.Bool("memory", false, "use in-memory storage")
)

func main() {
	flag.Parse()

	var taskRepository task.Repository
	if memory {
		taskRepository = internalTask.NewMemoryRepository()
	} else {
		taskRepository = internalTask.NewMemoryRepository()
	}

	taskService := internalTask.NewService(taskRepository)
	taskApi := taskapi.New(taskService)

	handler := NewHandler(taskApi)
	server := NewServer(port, handler)

	go server.Start()

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received signal: %v", (<-stopChannel).String())

	server.Stop()
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
