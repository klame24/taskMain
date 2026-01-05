package main

import (
	"context"
	"fmt"
	"net/http"
	"taskMain/internal/database"
	"taskMain/internal/handlers"
	"taskMain/internal/repositories"
	"taskMain/internal/routes"
	"taskMain/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Can't found .env file")
	}

	ctx := context.Background()

	conn, err := database.ConnectDB(ctx)
	if err != nil {
		fmt.Println("Can't connect to DB!")
	}

	fmt.Println("Successfully connected to DB!")

	// инициализация репозиториев
	userR := repositories.NewUserRepository(conn)
	projectR := repositories.NewProjectRepository(conn)

	// инициализация сервисов
	userS := services.NewUserService(userR)
	projectS := services.NewProjectService(projectR)

	// инициализация хендлеров
	userH := handlers.NewUserHandler(userS)
	projectH := handlers.NewProjectHandlers(projectS)

	// инициалазия сервера
	mux := http.NewServeMux()

	userRoutes := routes.SetupUserRoutes(userH)
	projectRoutes := routes.SetupProjectRoutes(projectH)

	mux.Handle("/users/", http.StripPrefix("/users", userRoutes))
	mux.Handle("/projects/", http.StripPrefix("/projects", projectRoutes))

	err = http.ListenAndServe(":5050", mux)
	if err != nil {
		panic(err)
	}
}
