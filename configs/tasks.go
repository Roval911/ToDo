package configs

import (
	"database/sql"
	"log"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      int       `json:"user"`
	Createdat   time.Time `json:"createdat"`
}

func GetTaskByIDAndOwner(taskID, userID int) (*Task, error) {
	var task Task
	query := `SELECT id, title, description, user_id, createdat
			  FROM tasks WHERE id = $1 AND user_id = $2`
	err := db.QueryRow(query, taskID, userID).Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.Createdat, &task.Completed)
	if err == sql.ErrNoRows {
		log.Printf("Задача не найдена или не принадлежит пользователю: %v", err)
		return nil, nil
	} else if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *Task) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2 ) RETURNING id`
	err := db.QueryRow(query, task.Title, task.Description).Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		log.Printf("Ошибка при добавлении задачи: %v", err)
		return err
	}
	return nil
}

func GetAllTasks() ([]Task, error) {
	query := `SELECT id, title, description, completed, createdat FROM tasks`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		// Порядок должен соответствовать структуре таблицы
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.Createdat)
		if err != nil {
			log.Printf("Ошибка чтения строки: %v", err)
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id int) (*Task, error) {
	query := `SELECT id, title, description FROM tasks WHERE id = $1`
	task := &Task{}
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description)

	if err == sql.ErrNoRows {
		log.Printf("Задача с ID %d не найдена", id) // Логируем отсутствие задачи
		return nil, nil
	} else if err != nil {
		log.Printf("Ошибка при получении задачи: %v", err)
		return nil, err
	}

	log.Printf("Найдена задача: %+v", task) // Логируем найденную задачу
	return task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE tasks SET title = $1, description = $2 WHERE id = $3`
	_, err := db.Exec(query, task.Title, task.Description, task.ID)
	if err != nil {
		log.Printf("Ошибка при обновлении задачи: %v", err)
		return err
	}
	return nil
}

func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Ошибка при удалении задачи: %v", err)
		return err
	}
	return nil
}
