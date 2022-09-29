package task

import (
	"time"
)

type CompleteStatus int

const (
	All CompleteStatus = iota
	Complete
	NotComplete
)

type (
	Task struct {
		ID         string    `json:"id"`
		Message    *string   `json:"message"`
		IsComplete *bool     `json:"is_complete"`
		CreatedAt  time.Time `json:"created_at"`
	}

	CreateRequest struct {
		Message string `json:"message" validate:"required,min=3"`
	}

	UpdateRequest struct {
		Message    *string `json:"message" validate:"omitempty,min=3"`
		IsComplete *bool   `json:"is_complete"`
	}
)
