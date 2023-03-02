package repository

import (
	"context"
	"errors"
	"kanbanApp/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var taskGet []entity.Task

	err := r.db.Model(&entity.Task{}).Where("user_id = ?", id).Scan(&taskGet).Error
	if err != nil {
		return nil, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}

	return taskGet, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	err = r.db.Create(&task).Error
	if err != nil {
		return 0, ctx.Err()
	}

	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var taskByID entity.Task
	err := r.db.Model(&entity.Task{}).Where("id = ?", id).Scan(&taskByID).Error
	if err != nil {
		return entity.Task{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Task{}, nil
	}

	return taskByID, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var res []entity.Task
	err := r.db.Model(&entity.Task{}).Where("category_id = ?", catId).Scan(&res).Error
	if err != nil {
		return nil, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}

	return res, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := r.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(task).Error
	if err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	if err != nil {
		return ctx.Err()
	}

	return nil // TODO: replace this
}
