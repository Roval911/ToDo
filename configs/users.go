package configs

import (
	"ToDo/internal/middleware"
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	passwordHash, err := middleware.HashPassword(user.Password)
	if err != nil {
		log.Printf("Ошибка при хэшировании пароля: %v", err)
		return fmt.Errorf("could not hash password: %w", err)
	}

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, user.Name, user.Email, passwordHash).Scan(&user.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении пользователя: %v", err)
		return err
	}
	return nil
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, password FROM users WHERE name = $1`
	err := db.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, err
	}
	return user, err
}

func GetUserByID(id int) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		log.Printf("Ошибка при обновлении пользователя: %v", err)
		return err
	}
	return nil
}

// DeleteUser удаляет пользователя по ID
func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Ошибка при удалении пользователя: %v", err)
		return err
	}
	return nil
}
