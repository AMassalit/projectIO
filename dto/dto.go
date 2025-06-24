package dto

import "time"

type Task struct {
	ID        int
	Status    string
	CreatedAt time.Time
	Expiry    time.Time
}

type TaskResponse struct {
	ID        int
	Status    string
	CreatedAt string
	Duration  string
}
