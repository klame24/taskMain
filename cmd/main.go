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

	// инициализация сервисов
	userS := services.NewUserService(userR)

	// инициализация хендлеров
	userH := handlers.NewUserHandler(userS)

	// инициализация роутеров
	userRoutes := routes.SetupRoutes(userH)

	http.ListenAndServe(":5050", userRoutes)
}
