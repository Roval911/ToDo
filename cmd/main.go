package main

import (
	"ToDo/configs"
	_ "ToDo/docs" // Импорт сгенерированной документации Swagger
	"ToDo/internal/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files" // Импортируйте с алиасом для ясности
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title ToDo API
// @version 1.0
// @description API для управления пользователями и задачами
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

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
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("api/users", controllers.CreateUserHandler)
	router.GET("api/users/:id", controllers.GetUserHandler)
	router.PUT("api/users/:id", controllers.UpdateUserHandler)
	router.DELETE("api/users/:id", controllers.DeleteUserHandler)
	router.POST("api/login", controllers.LoginHandler)
	router.POST("api/logout", controllers.LogoutHandler)

	router.POST("api/task", controllers.CreaateTaskHandler)
	router.GET("api/tasks/:id", controllers.GetTaskHandler)
	router.GET("api/tasks", controllers.GetAllTasksHandler)
	router.PUT("api/tasks/:id", controllers.UpdateTaskHandler)
	router.DELETE("api/tasks/:id", controllers.DeleteTaskHandler)

	router.Run(os.Getenv("PORT"))
}
