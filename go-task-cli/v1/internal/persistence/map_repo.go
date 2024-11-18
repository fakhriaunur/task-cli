package persistence

import (
	"errors"
	"fmt"
	"sync"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

type MapRepo struct {
	mu *sync.RWMutex
	MapRepoStructure
}

type MapRepoStructure struct {
	tasks map[int]tt.Task
}

var Tasks map[int]tt.Task

var _ tt.TaskRepoInMemoryPort[map[int]tt.Task] = (*MapRepo)(nil)

func NewMapRepo() (tt.TaskRepoInMemoryPort[map[int]tt.Task], error) {
	return &MapRepo{
		mu: &sync.RWMutex{},
		MapRepoStructure: MapRepoStructure{
			tasks: make(map[int]tt.Task),
		},
	}, nil
}

// Write implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Write(tasks map[int]tt.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks = tasks

	return nil
}

// Load implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Load() (map[int]tt.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.tasks, nil
}

func (r *MapRepo) Reset() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks = nil

	return nil
}

// TODO: refactor to a new file, jason_repo_task.go

// Add implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Add(desc string) (tt.Task, error) {
	fmt.Println("begin add repo")

	tasks, err := r.Load()
	if err != nil {
		return tt.Task{}, err
	}

	newID := len(tasks) + 1
	newTask := tt.Task{
		ID:          newID,
		Status:      tt.StatusTodo,
		Description: desc,
		CreatedAt:   tt.TimeNowLocal,
		UpdatedAt:   tt.TimeZero,
	}

	if err := r.Write(tasks); err != nil {
		return tt.Task{}, err
	}

	fmt.Println("end add repo")
	return newTask, nil
}

// Delete implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Delete(id int) error {
	_, ok := r.tasks[id]
	if !ok {
		return errors.New("task not existed")
	}

	delete(r.tasks, id)

	return nil
}

func (r *MapRepo) ListAll() ([]tt.Task, error) {
	tasks := make([]tt.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// List implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) List(status string) ([]tt.Task, error) {
	tasks := make([]tt.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// Mark implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Mark(id int, status string) error {
	task, ok := r.tasks[id]
	if !ok {
		return errors.New("")
	}

	task.Status = status
	r.tasks[id] = task

	return nil
}

// Reset implements tasktracker.TaskRepoInMemory.

// Update implements tasktracker.TaskRepoInMemory.
func (r *MapRepo) Update(id int, status string) error {
	task, ok := r.tasks[id]
	if !ok {
		return errors.New("")
	}

	task.Status = status
	r.tasks[id] = task

	return nil
}
