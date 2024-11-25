package main

import (
	"ToDo/configs"
	"ToDo/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Инициализация базы данных
	configs.InitDb()

	// Запуск миграций
	configs.RunMigrations()
}

func main() {
	router := gin.Default()

	router.POST("api/user", controllers.CreateUserHandler)
	router.GET("api/user/:id", controllers.GetUserHandler)
	router.PUT("api/user/:id", controllers.UpdateUserHandler)
	router.DELETE("api/user/:id", controllers.DeleteUserHandler)

	router.Run(":8080")
}
