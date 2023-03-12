package usecase

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"time"
)

type taskUsecase struct {
	taskRepository model.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository model.TaskRepository, timeout time.Duration) model.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *model.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, userID)
}
