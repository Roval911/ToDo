package main

import (
	"ToDo/configs"
	"ToDo/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	router.POST("api/task", controllers.CreaateTaskHandler)
	router.GET("api/task/:id", controllers.GetTaskHandler)
	router.GET("api/tasks", controllers.GetAllTasksHandler)
	router.PUT("api/task/:id", controllers.UpdateTaskHandler)
	router.DELETE("api/task/:id", controllers.DeleteTaskHandler)

	router.Run(os.Getenv("PORT"))
}
