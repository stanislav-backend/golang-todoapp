package tasks_postgres_repository

import (
	"time"

	"github.com/stanislav-backend/golang-todoapp/internal/core/domain"
)

type TaskModel struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserId int
}

func taskDomainsFromModels(tasksModels []TaskModel) []domain.Task{
	domains := make([]domain.Task, len(tasksModels))

	for i, model := range tasksModels{
		domains[i] = TaskDomainFromModel(model)
	}
	return domains
}


func TaskDomainFromModel(taskModel TaskModel) domain.Task{
	return domain.NewTask(taskModel.ID, 
		taskModel.Version, 
		taskModel.Title, 
		taskModel.Description, 
		taskModel.Completed, 
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserId,)
}