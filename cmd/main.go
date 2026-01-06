package main

import (
	"context"
	"fmt"
	"net/http"
	"taskMain/internal/auth/jwt"
	"taskMain/internal/database/mongo"
	"taskMain/internal/database/postgres"
	"taskMain/internal/handlers"
	"taskMain/internal/repositories"
	"taskMain/internal/routes"
	"taskMain/internal/services"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Can't found .env file")
	}

	ctx := context.Background()

	conn, err := postgres.ConnectPostgresDB(ctx)
	if err != nil {
		fmt.Println("Can't connect to Postgres DB!")
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres DB!")

	client, err := mongo.ConnectMongoDB(ctx)
	if err != nil {
		fmt.Println("Can't connect to Mongo DB!")
		panic(err)
	}

	fmt.Println("Successfully connected to Mongo DB!")

	// jwt
	jwtSecret := "secret-key"

	accessTokenExp := 24 * time.Hour
	refreshTokenExp := 7 * 24 * time.Hour 

	jwtManager := jwt.NewManager(jwtSecret, accessTokenExp)

	// инициализация репозиториев
	userR := repositories.NewUserRepository(conn)
	projectR := repositories.NewProjectRepository(conn)
	taskR := repositories.NewTaskRepository(conn)
	tokenR := jwt.NewTokenRepository(client)

	// инициализация сервисов
	userS := services.NewUserService(userR)
	projectS := services.NewProjectService(projectR)
	taskS := services.NewTaskService(taskR)
	authS := services.NewAuthService(userR, tokenR, jwtManager, refreshTokenExp)

	// инициализация хендлеров
	userH := handlers.NewUserHandler(userS)
	projectH := handlers.NewProjectHandlers(projectS)
	taskH := handlers.NewTaskHandler(taskS)
	authH := handlers.NewAuthHandler(authS)

	// инициалазия сервера
	mux := http.NewServeMux()

	userRoutes := routes.SetupUserRoutes(userH)
	projectRoutes := routes.SetupProjectRoutes(projectH, taskH)
	taskRoutes := routes.SetupTaskRoutes(taskH)
	authRoutes := routes.SetupAuthRoutes(authH)

	mux.Handle("/users/", http.StripPrefix("/users", userRoutes))
	mux.Handle("/projects/", http.StripPrefix("/projects", projectRoutes))
	mux.Handle("/tasks/", http.StripPrefix("/tasks", taskRoutes))
	mux.Handle("/auth/", http.StripPrefix("/auth", authRoutes))

	err = http.ListenAndServe(":5050", mux)
	if err != nil {
		panic(err)
	}
}
