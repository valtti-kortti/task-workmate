package repository

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ProcessTime int64     `json:"process_time"`
}
