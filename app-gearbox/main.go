package main

import (
	"appgearbox/incoming/taskapi"
	"appgearbox/internal/port"
	"appgearbox/internal/service"
	"appgearbox/outgoing/taskrepo"
	"log"
	"os/signal"
	"strconv"
	"syscall"

	"fmt"
	"os"
	"time"

	"github.com/gogearbox/gearbox"
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

	go handler.Start(fmt.Sprintf(":%d", serverPortInt))

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("received signal: %v\n", (<-stopChannel).String())

	handler.Stop()
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

func NewHandler(taskApi *taskapi.API) gearbox.Gearbox {
	router := gearbox.New()

	taskApi.RegisterRoutes(router)

	return router
}
