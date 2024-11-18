package persistence

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

type JSONRepo struct {
	path string
	mu   *sync.RWMutex
}

var _ tt.TaskRepoInMemoryPort[JSONStructure] = (*JSONRepo)(nil)

type JSONStructure struct {
	Tasks map[int]tt.Task
}

func NewJSONRepo(path string) (tt.TaskRepoInMemoryPort[JSONStructure], error) {
	r := &JSONRepo{
		path: path,
		mu:   &sync.RWMutex{},
	}

	err := r.ensureRepo()

	return r, err
}

func (r *JSONRepo) createRepo() error {
	js := JSONStructure{
		// Tasks: make(map[int]tt.Task),
		Tasks: map[int]tt.Task{},
	}

	return r.Write(js)
}

func (r *JSONRepo) ensureRepo() error {
	_, err := os.ReadFile(r.path)
	if errors.Is(err, os.ErrNotExist) {
		return r.createRepo()
	}

	return err
}

// Write implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) write(js JSONStructure) error {
	log.Println("yo")
	r.mu.Lock()
	defer r.mu.Unlock()

	dat, err := json.MarshalIndent(js, "", "\t")
	if err != nil {
		return err
	}

	var buffer *bytes.Buffer
	if err := gob.NewEncoder(buffer).Encode(js); err != nil {

	}

	if err := os.WriteFile(r.path, dat, 0644); err != nil {
		return err
	}

	return nil

}
func (r *JSONRepo) Write(js JSONStructure) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.OpenFile(
		r.path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(js); err != nil {
		return err
	}

	return nil
}

func (r *JSONRepo) Load() (JSONStructure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	file, err := os.OpenFile(
		r.path,
		os.O_RDONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		return JSONStructure{}, err
	}

	var js JSONStructure
	if err := json.NewDecoder(file).Decode(&js); err != nil && err.Error() != "EOF" {
		return JSONStructure{}, err
	}

	return js, nil
}

// Load implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) load() (JSONStructure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var js JSONStructure
	dat, err := os.ReadFile(r.path)
	if err != nil {
		fmt.Println("no file")
		return JSONStructure{}, err
	}

	buffer := bytes.NewBuffer(dat)

	if err := json.NewDecoder(buffer).Decode(&js); err != nil {
		return js, err
	}

	if err := json.Unmarshal(dat, &js); err != nil {
		return JSONStructure{}, err
	}

	return js, nil
}

func (r *JSONRepo) Reset() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := os.Remove(r.path); err != nil {
		return err
	}

	_, err := os.Create(r.path)
	if err != nil {
		return err
	}

	// // var js JSONStructure
	// if err := json.NewEncoder(file).Encode(&js); err != nil {
	// 	return err
	// }

	return nil
}

// TODO: refactor to a new file, json_repo_task.go

// Add implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) Add(desc string) (tt.Task, error) {
	js, err := r.Load()
	if err != nil {
		fmt.Println("error loading")
		return tt.Task{}, err
	}

	newID := len(js.Tasks) + 1
	newTask := tt.Task{
		ID:          newID,
		Status:      tt.StatusTodo,
		Description: desc,
		CreatedAt:   tt.TimeNowLocal,
		UpdatedAt:   tt.TimeZero,
	}
	js.Tasks[newID] = newTask

	// fmt.Println(newTask)
	if err := r.Write(js); err != nil {
		fmt.Println("error writing")
		return tt.Task{}, err
	}

	return newTask, nil
}

// Delete implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) Delete(id int) error {
	js, err := r.Load()
	if err != nil {
		return err
	}

	_, ok := js.Tasks[id]
	if !ok {
		return errors.New("task not existed")
	}

	delete(js.Tasks, id)

	return nil
}

func (r *JSONRepo) ListAll() ([]tt.Task, error) {
	js, err := r.Load()
	if err != nil {
		return nil, err
	}

	tasks := make([]tt.Task, 0, len(js.Tasks))
	for _, task := range js.Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// List implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) List(status string) ([]tt.Task, error) {
	js, err := r.Load()
	if err != nil {
		return nil, err
	}

	tasks := make([]tt.Task, 0, len(js.Tasks))
	for _, task := range js.Tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// Mark implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) Mark(id int, status string) error {
	js, err := r.Load()
	if err != nil {
		return err
	}

	task, ok := js.Tasks[id]
	if !ok {
		return errors.New("")
	}

	task.Status = status
	js.Tasks[id] = task

	return nil
}

// Reset implements tasktracker.TaskRepoInMemory.

// Update implements tasktracker.TaskRepoInMemory.
func (r *JSONRepo) Update(id int, status string) error {
	js, err := r.Load()
	if err != nil {
		return err
	}

	task, ok := js.Tasks[id]
	if !ok {
		return errors.New("task not existed")
	}

	task.Status = status
	js.Tasks[id] = task

	return nil
}
