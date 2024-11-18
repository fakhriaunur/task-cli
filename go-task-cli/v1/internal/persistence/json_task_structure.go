package persistence

import (
	"errors"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

// var _ tt.TaskRepoInMemoryPort = (*JSONTaskStructure)(nil)

type JSONTaskStructure struct {
	Tasks map[int]tt.Task
}

func (repo *JSONTaskRepo) Add(desc string) (tt.Task, error) {
	jsonTaskRepo, err := repo.Load()
	if err != nil {
		return tt.Task{}, err
	}

	id := len(jsonTaskRepo.Tasks) + 1
	newTask := tt.Task{
		ID:          id,
		Description: desc,
		Status:      tt.StatusTodo,
		CreatedAt:   tt.TimeNowLocal,
		UpdatedAt:   tt.TimeZero,
	}
	jsonTaskRepo.Tasks[id] = newTask

	err = repo.Write(jsonTaskRepo)
	if err != nil {
		return tt.Task{}, err
	}

	return newTask, nil
}

func (repo *JSONTaskRepo) Delete(id int) error {
	jsonTaskRepo, err := repo.Load()
	if err != nil {
		return err
	}

	_, ok := jsonTaskRepo.Tasks[id]
	if !ok {
		return errors.New("")
	}

	delete(jsonTaskRepo.Tasks, id)

	if err := repo.Write(jsonTaskRepo); err != nil {
		return err
	}

	return nil
}

func (repo *JSONTaskRepo) Update(id int, desc string) error {
	return nil
}

func (repo *JSONTaskRepo) Mark(id int, status string) error {
	return nil
}

func (repo *JSONTaskRepo) List(status string) ([]tt.Task, error) {
	return nil, nil
}

func (repo *JSONTaskRepo) ListAll() ([]tt.Task, error) {
	return nil, nil
}
