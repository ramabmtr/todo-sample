package task

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/todo-sample/util"
)

type UseCaseIFace interface {
	Create(ctx context.Context, req *CreateRequest) (*Task, error)
	Get(ctx context.Context, limit, page int, completeStatus CompleteStatus) ([]*Task, int, error)
	GetByID(ctx context.Context, id string) (*Task, error)
	UpdateByID(ctx context.Context, id string, req *UpdateRequest) (*Task, error)
	DeleteByID(ctx context.Context, id string) error
}

type UseCase struct {
	repo RepoIFace
}

func NewUseCase(repo RepoIFace) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, req *CreateRequest) (*Task, error) {
	task := &Task{
		ID:         util.GenerateRandom(6),
		Message:    &req.Message,
		IsComplete: util.GetBoolPointer(false),
		CreatedAt:  time.Now().UTC(),
	}
	err := u.repo.CreateTask(ctx, task)
	return task, err
}

func (u *UseCase) Get(ctx context.Context, limit, page int, completeStatus CompleteStatus) ([]*Task, int, error) {
	return u.repo.GetTasks(ctx, limit, page, completeStatus)
}

func (u *UseCase) GetByID(ctx context.Context, id string) (*Task, error) {
	return u.repo.GetTaskByID(ctx, id)
}

func (u *UseCase) UpdateByID(ctx context.Context, id string, req *UpdateRequest) (*Task, error) {
	oldTask, err := u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if oldTask == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "data not found")
	}

	msg := oldTask.Message
	if req.Message != nil {
		msg = req.Message
	}

	isComplete := oldTask.IsComplete
	if req.IsComplete != nil {
		isComplete = req.IsComplete
	}

	task := &Task{
		ID:         id,
		Message:    msg,
		IsComplete: isComplete,
		CreatedAt:  oldTask.CreatedAt,
	}
	err = u.repo.UpdateTask(ctx, task)
	return task, err
}

func (u *UseCase) DeleteByID(ctx context.Context, id string) error {
	return u.repo.DeleteTaskByID(ctx, id)
}
