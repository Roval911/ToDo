package models

import "time"

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Tasks struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Userid      int       `json:"user"`
	Created_at  time.Time `json:"created_at"`
}
