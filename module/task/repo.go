package task

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/ramabmtr/todo-sample/util"
)

type RepoIFace interface {
	GetTasks(ctx context.Context, limit, page int, completeStatus CompleteStatus) (tasks []*Task, total int, err error)
	GetTaskByID(ctx context.Context, id string) (*Task, error)
	CreateTask(ctx context.Context, t *Task) error
	UpdateTask(ctx context.Context, t *Task) error
	DeleteTaskByID(ctx context.Context, id string) error
}

type Repo struct {
	c *cache.Cache
}

func NewRepo(c *cache.Cache) *Repo {
	return &Repo{
		c: c,
	}
}

func (r *Repo) GetTasks(ctx context.Context, limit, page int, completeStatus CompleteStatus) (tasks []*Task, total int, err error) {
	tasks = make([]*Task, 0)

	data := r.c.Items()
	for _, item := range data {
		if task, ok := item.Object.(*Task); ok {
			switch completeStatus {
			case Complete:
				if util.GetBool(task.IsComplete) {
					tasks = append(tasks, task)
				}
			case NotComplete:
				if !util.GetBool(task.IsComplete) {
					tasks = append(tasks, task)
				}
			default:
				tasks = append(tasks, task)
			}
		}
	}

	total = len(tasks)
	if page+limit == 0 {
		return tasks, total, nil
	}

	if page > total {
		return []*Task{}, total, nil
	}

	if page+limit > total {
		return tasks[page:total], total, nil
	}

	return tasks[page : page+limit], total, nil
}

func (r *Repo) GetTaskByID(ctx context.Context, id string) (*Task, error) {
	val, ok := r.c.Get(id)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusNotFound, "task not found")
	}

	task, ok := val.(*Task)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "task is broken")
	}

	return task, nil
}

func (r *Repo) CreateTask(ctx context.Context, t *Task) error {
	r.c.Set(t.ID, t, cache.NoExpiration)
	return nil
}

func (r *Repo) UpdateTask(ctx context.Context, t *Task) error {
	r.c.Set(t.ID, t, cache.NoExpiration)
	return nil
}

func (r *Repo) DeleteTaskByID(ctx context.Context, id string) error {
	r.c.Delete(id)
	return nil
}
