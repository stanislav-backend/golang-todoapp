package tasks_transport_http

import (
	"time"

	"github.com/stanislav-backend/golang-todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id" example:"15"`
	Version      int        `json:"version" example:"3"`
	Title        string     `json:"title" example:"Домашка"`
	Description  *string    `json:"description" example:"Сделать до четверга домашнее задание по математике"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserId int        `json:"author_user_id" example:"5"`
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