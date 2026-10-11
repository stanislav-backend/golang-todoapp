package tasks_transport_http

import (
	"time"

	"github.com/stanislav-backend/golang-todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorUserId int `json:"author_user_id"`
}

func taskDTOFromDomain(taskDomain domain.Task) TaskDTOResponse{
	return TaskDTOResponse{
		ID: taskDomain.ID,
		Version: taskDomain.Version,
		Title: taskDomain.Title,
		Description: taskDomain.Description,
		Completed: taskDomain.Completed,
		CreatedAt: taskDomain.CreatedAt,
		CompletedAt: taskDomain.CompletedAt,

		AuthorUserId: taskDomain.AuthorUserId,
	}
}

func taskDTOsFromDomains(tasks []domain.Task) []TaskDTOResponse{
	dtos := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks{
		dtos[i] = taskDTOFromDomain(task)
	}

	return dtos
}