package controllers

import (
	"ToDo/configs"
	"ToDo/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUserHandler создает нового пользователя
// @Summary Создать пользователя
// @Description Создает нового пользователя и сохраняет его в базе данных
// @Tags Пользователи
// @Param user body configs.User true "Данные пользователя"
// @Success 201 {object} configs.User
// @Failure 400 {object} map[string]string "Неверный формат данных"
// @Failure 500 {object} map[string]string "Не удалось создать пользователя"
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var user configs.User

	// Чтение JSON из тела запроса
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := configs.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserHandler возвращает информацию о пользователе по его ID
// @Summary Получить пользователя
// @Description Возвращает информацию о пользователе по его ID
// @Tags Пользователи
// @Param id path int true "ID пользователя"
// @Success 200 {object} configs.User
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Ошибка базы данных"
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	user, err := configs.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler обновляет данные пользователя по ID
// @Summary Обновить пользователя
// @Description Обновляет данные пользователя по его ID
// @Tags Пользователи
// @Param id path int true "ID пользователя"
// @Param user body configs.User true "Данные пользователя"
// @Success 200 {object} configs.User
// @Failure 400 {object} map[string]string "Неверный ID или формат данных"
// @Failure 500 {object} map[string]string "Не удалось обновить пользователя"
// @Router /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var user configs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	user.ID = id

	if err := configs.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить пользователя"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler удаляет пользователя по его ID
// @Summary Удалить пользователя
// @Description Удаляет пользователя из базы данных по его ID
// @Tags Пользователи
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]string "Пользователь удален"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 500 {object} map[string]string "Не удалось удалить пользователя"
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if err := configs.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален"})
}

// loginHandler авторизует пользователя и возвращает JWT
// @Summary Авторизация
// @Description Авторизация пользователя по email и паролю
// @Tags Авторизация
// @Param credentials body map[string]string true "Email и пароль"
// @Success 200 {object} map[string]string "JWT токен"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 401 {object} map[string]string "Неверный email или пароль"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /login [post]
func loginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Ищем пользователя по email
	user, err := configs.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if err == nil || !middleware.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Генерируем токен с именем пользователя
	token, err := middleware.GenerateJWT(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Возвращаем токен клиенту
	c.JSON(http.StatusOK, gin.H{"token": token})
}
