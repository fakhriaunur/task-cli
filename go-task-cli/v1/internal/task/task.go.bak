package task

import (
	"errors"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

var localNow = time.Now().Local()

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// var tasks []Task

// func GetTasks() []Task {
// 	return tasks
// }

func Add(desc string) (Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   localNow,
		UpdatedAt:   localNow,
	}

	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func Update(id int, newDesc string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = localNow

			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func Delete(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return nil
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			return SaveTasks(tasks)
		}
	}

	return errors.New("couldnt find task")
}

// TODO: IMPLEMENT LIST, NO ARGS
func list(status string) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []Task
	switch status {
	case "":
		return tasks, nil
	case StatusDone, StatusTodo, StatusInProgress:

		// filteredTasks := []Task{}
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	default:
		return nil, errors.New("unknown status")
	}

	return filteredTasks, nil
}

func ListAll() ([]Task, error) {
	return list("")
}

func ListByStatus(status string) ([]Task, error) {
	return list(status)
}

// TODO: Implement Mark
func MarkInProgress(id int) error {
	return mark(id, StatusInProgress)
}

func MarkDone(id int) error {
	return mark(id, StatusDone)
}

func mark(id int, status string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return nil
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status

			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
