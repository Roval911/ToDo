package controllers

import (
	"ToDo/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreaateTaskHandler(c *gin.Context) {
	var task configs.Task
	err := c.ShouldBindBodyWithJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	configs.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать задачу"})
		return
	}
}

func GetTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	task, err := configs.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetAllTasksHandler(c *gin.Context) {
	tasks, err := configs.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}
	if tasks == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Список задач пуст"})
		return
	}
	c.JSON(http.StatusOK, tasks)

}

func UpdateTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var task configs.Task
	err = c.ShouldBindBodyWithJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	task.ID = id

	err = configs.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
		return
	}
}

func DeleteTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if err := configs.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить задачу"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
