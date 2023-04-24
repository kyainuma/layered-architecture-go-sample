package infra

import (
	"sample/domain/model"
	"sample/domain/repository"

	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	Coon *gorm.DB
}

func NewTaskRepository(coon *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Coon: coon}
}

func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Coon.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Coon.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Coon.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Coon.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}