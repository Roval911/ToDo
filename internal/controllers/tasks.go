package controllers

import (
	"ToDo/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreaateTaskHandler создает новую задачу
// @Summary Создать задачу
// @Description Создает новую задачу для авторизованного пользователя
// @Tags Задачи
// @Param task body configs.Task true "Данные задачи"
// @Success 200 {object} configs.Task
// @Failure 400 {object} map[string]string "Неверный формат данных"
// @Failure 401 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Не удалось создать задачу"
// @Router /tasks [post]
func CreaateTaskHandler(c *gin.Context) {
	//username, _ := c.Get("name")
	// Получаем ID пользователя из базы данных по имени (или email, если нужно)
	//user, err := configs.GetUserByName(username.(string)) // Предполагается, что GetUserByName возвращает пользователя по имени
	//if err != nil || user == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
	//	return
	//}
	var task configs.Task
	err := c.ShouldBindBodyWithJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	//task.UserID = user.ID

	err = configs.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать задачу"})
		return
	}
}

// GetTaskHandler возвращает информацию о задаче по её ID
// @Summary Получить задачу
// @Description Возвращает данные задачи по её ID
// @Tags Задачи
// @Param id path int true "ID задачи"
// @Success 200 {object} configs.Task
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Задача не найдена"
// @Failure 500 {object} map[string]string "Ошибка базы данных"
// @Router /tasks/{id} [get]
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

// GetAllTasksHandler возвращает список всех задач
// @Summary Получить список задач
// @Description Возвращает список всех задач из базы данных
// @Tags Задачи
// @Success 200 {array} configs.Task
// @Failure 404 {object} map[string]string "Список задач пуст"
// @Failure 500 {object} map[string]string "Ошибка базы данных"
// @Router /tasks [get]
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

// UpdateTaskHandler обновляет задачу
// @Summary Обновить задачу
// @Description Обновляет данные задачи по её ID
// @Tags Задачи
// @Param id path int true "ID задачи"
// @Param task body configs.Task true "Данные задачи"
// @Success 200 {object} configs.Task
// @Failure 400 {object} map[string]string "Неверный ID или формат данных"
// @Failure 401 {object} map[string]string "Пользователь не найден"
// @Failure 403 {object} map[string]string "Вы не можете обновить эту задачу"
// @Failure 500 {object} map[string]string "Не удалось обновить задачу"
// @Router /tasks/{id} [put]
func UpdateTaskHandler(c *gin.Context) {
	//username, _ := c.Get("name")
	//user, err := configs.GetUserByName(username.(string))
	//if err != nil || user == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
	//	return
	//}
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
	//task.UserID = user.ID

	//existingTask, _ := configs.GetTaskByIDAndOwner(id, user.ID)
	//if existingTask == nil {
	//	c.JSON(http.StatusForbidden, gin.H{"error": "Вы не можете обновить эту задачу"})
	//	return
	//}

	err = configs.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
		return
	}
}

// DeleteTaskHandler удаляет задачу
// @Summary Удалить задачу
// @Description Удаляет задачу по её ID, если пользователь является её владельцем
// @Tags Задачи
// @Param id path int true "ID задачи"
// @Success 200 {object} map[string]string "Задача удалена"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 401 {object} map[string]string "Пользователь не найден"
// @Failure 403 {object} map[string]string "Вы не можете удалить эту задачу"
// @Failure 500 {object} map[string]string "Не удалось удалить задачу"
// @Router /tasks/{id} [delete]
func DeleteTaskHandler(c *gin.Context) {
	//username, _ := c.Get("name")
	//user, err := configs.GetUserByName(username.(string))
	//if err != nil || user == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
	//	return
	//}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	//existingTask, _ := configs.GetTaskByIDAndOwner(id, user.ID)
	//if existingTask == nil {
	//	c.JSON(http.StatusForbidden, gin.H{"error": "Вы не можете удалить эту задачу"})
	//	return
	//}

	if err := configs.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить задачу"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
