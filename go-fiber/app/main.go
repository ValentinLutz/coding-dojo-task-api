package main

import (
	"appfiber/incoming/openapi"
	"appfiber/incoming/taskapi"
	"appfiber/internal/port"
	"appfiber/internal/service"
	"appfiber/outgoing/taskrepo"
	"context"
	"log"
	"os/signal"
	"strconv"
	"syscall"

	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
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
	openApi := openapi.New()

	app := newApp(taskApi, openApi)
	go app.Listen(fmt.Sprintf(":%d", serverPortInt))

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received signal: %v\n", (<-stopChannel).String())

	app.Shutdown()
}

func NewDatabase() *pgxpool.Pool {
	databaseHost := os.Getenv("POSTGRES_HOST")
	databasePort := os.Getenv("POSTGRES_PORT")
	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")

	psqlInfo := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable pool_max_conns=100",
		databaseHost, databasePort, databaseUser, databasePassword, databaseName,
	)

	pgxConfig, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		log.Fatalf("failed to parse pgx config: %v\n", err)
	}
	db, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	initTasksTable(db)

	return db
}

func initTasksTable(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(), "DROP TABLE IF EXISTS public.tasks")
	if err != nil {
		log.Fatalf("failed to drop table public.tasks: %v\n", err)
	}
	log.Println("dropped table public.tasks")

	_, err = db.Exec(
		context.Background(),
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

func newApp(taskApi *taskapi.API, openApi *openapi.API) *fiber.App {
	app := fiber.New()

	taskApi.RegisterRoutes(app)
	openApi.RegisterRoutes(app)

	return app
}
