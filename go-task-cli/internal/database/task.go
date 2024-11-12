package database

import "time"

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var Tasks []Task

func Add(desc string) (*Task, error) {
	task := &Task{
		ID:          globalID + 1,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Time{},
	}

	return task, nil
}

func (t *Task) Update(id int, desc string) error {
	t.Description = desc

	return nil
}
