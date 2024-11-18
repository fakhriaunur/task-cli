package persistence

import (
	"encoding/json"
	"os"
	"sync"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

type JSONTaskRepo struct {
	path string
	mu   *sync.RWMutex
}

func NewJSONTaskRepo(path string) (*JSONTaskRepo, error) {
	repo := &JSONTaskRepo{
		path: path,
		mu:   &sync.RWMutex{},
	}

	return repo, nil
}

var _ tt.TaskRepoInMemoryPort[JSONTaskStructure] = (*JSONTaskRepo)(nil)

func (repo *JSONTaskRepo) Write(jts JSONTaskStructure) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	file, err := os.OpenFile(
		repo.path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(jts); err != nil {
		return err
	}

	return nil
}

func (repo *JSONTaskRepo) Load() (JSONTaskStructure, error) {
	repo.mu.RLock()
	defer repo.mu.Unlock()

	var jts JSONTaskStructure
	file, err := os.OpenFile(
		repo.path,
		os.O_RDONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		return jts, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jts); err != nil {
		return jts, err
	}

	return jts, nil
}

func (repo *JSONTaskRepo) Reset() error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if err := os.Remove(repo.path); err != nil {
		return err
	}

	return nil
}
