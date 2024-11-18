package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/fakhriaunur/task-cli/internal/database"
)

type TaskStore interface {
	database.Database
	GetTasks() map[int]Task
}

type TaskManager struct {
	db TaskStore
}

func NewTaskManager(db TaskStore) *TaskManager {
	return &TaskManager{
		db: db,
	}
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
	StatusUnknown    = "unknown status"
)

var (
	TimeNowLocal = time.Now().Local()
	TimeZero     = time.Time{}
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TODO: Refactor the commands / query integreated with DB
func (t *TaskManager) Add(desc string) (Task, error) {
	dbStructure, err := t.db.LoadDB()
	if err != nil {
		return Task{}, err
	}

	id := len(dbStructure.Tasks) + 1
	task := Task{
		ID:          id,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   TimeNowLocal,
		UpdatedAt:   TimeZero,
	}
	dbStructure.Tasks[id] = task

	err = t.db.WriteDB(dbStructure)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t *TaskManager) Update(id int, desc string) error {
	dbStructure, err := t.db.LoadDB()
	if err != nil {
		return err
	}

	task, ok := dbStructure.Tasks[id]
	if !ok {
		return database.ErrNotExist
	}

	task.Description = desc
	task.UpdatedAt = TimeNowLocal
	dbStructure.Tasks[id] = task

	err = t.db.WriteDB(dbStructure)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskManager) Delete(id int) error {
	dbStructure, err := t.db.LoadDB()
	if err != nil {
		return err
	}

	_, ok := dbStructure.Tasks[id]
	if !ok {
		return database.ErrNotExist
	}

	delete(dbStructure.Tasks, id)

	err = t.db.WriteDB(dbStructure)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskManager) MarkInProgress(id int) error {
	fmt.Println("before hello")
	return t.mark(id, StatusInProgress)
}

func (t *TaskManager) MarkDone(id int) error {
	fmt.Println("before hello")
	return t.mark(id, StatusDone)
}

func (t *TaskManager) mark(id int, status string) error {
	fmt.Println("hello")
	dbStructure, err := t.db.LoadDB()
	if err != nil {
		return err
	}
	fmt.Println("loaded db")

	task, ok := dbStructure.Tasks[id]
	if !ok {
		return database.ErrNotExist
	}
	fmt.Println("task loaded")

	task.Status = status
	dbStructure.Tasks[id] = task

	err = t.db.WriteDB(dbStructure)
	if err != nil {
		return err
	}
	fmt.Println("written to db")

	return nil
}

func (t *TaskManager) ListAll() ([]Task, error) {
	return t.list("")
}
func (t *TaskManager) ListByStatus(status string) ([]Task, error) {
	return t.list(status)
}
func (t *TaskManager) list(status string) ([]Task, error) {
	dbStructure, err := t.db.LoadDB()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	switch status {
	case "":
		fmt.Printf("len begin: %d\n", len(dbStructure.Tasks))
		for _, task := range dbStructure.Tasks {
			tasks = append(tasks, task)
		}
		return tasks, nil

	case StatusDone, StatusInProgress, StatusTodo:
		fmt.Printf("len begin: %d\n", len(dbStructure.Tasks))
		for _, task := range dbStructure.Tasks {
			fmt.Printf("task.id: %d\ttask.status: %s\tstatus: %s\n", task.ID, task.Status, status)
			if task.Status == status {
				fmt.Printf("task.status: %s\tstatus: %s\n", task.Status, status)
				tasks = append(tasks, task)
				fmt.Printf("len end: %d\n", len(tasks))
			}
		}
	default:
		return nil, errors.New(StatusUnknown)
	}

	return tasks, nil
}
