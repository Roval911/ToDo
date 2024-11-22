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

	router.GET("api/tasks", controllers.GetTasks())
	router.POST("api/tasks", controllers.CreateTask())
	router.GET("api/tasks/:id", controllers.GetTaskByID())
	router.PUT("api/tasks/:id", controllers.UpdateTask())
	router.DELETE("api/tasks/:id", controllers.DeleteTask())

	router.Run(":8080")
}
