package task

import (
	"os"
	"sync"
)

var _ taskDAOPort = (*taskDAO)(nil)

type taskDAO struct {
	path   string
	mu     *sync.RWMutex
	mapper taskFileMapperPort
}

type taskStructure struct {
	Tasks map[int]task
}

func NewTaskDAO(path string, mapper taskFileMapperPort) taskDAOPort {
	return &taskDAO{
		path:   path,
		mu:     &sync.RWMutex{},
		mapper: mapper,
	}
}

// loadTaskStructure implements taskDAOPort.
func (t *taskDAO) loadTaskStructure() (taskStructure, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	file, err := os.OpenFile(t.path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return taskStructure{}, err
	}
	defer file.Close()

	return t.mapper.decode(file)
}

// saveTaskStructure implements taskDAOPort.
func (t *taskDAO) saveTaskStructure(ts taskStructure) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	file, err := os.OpenFile(
		t.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644,
	)
	if err != nil {
		return nil
	}

	err = t.mapper.encode(file, ts)
	if err != nil {
		return err
	}

	return nil
}
