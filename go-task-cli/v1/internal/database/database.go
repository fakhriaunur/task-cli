package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/fakhriaunur/task-cli/internal/task"
)

var (
	ErrNotExist     = errors.New("resource not found")
	ErrAlreadyExist = errors.New("resource already exist")
)

type Database interface {
	CreateDB() error
	EnsureDB() error
	WriteDB(DBStructure) error
	LoadDB() (DBStructure, error)
	ResetDB() error
}

type DB struct {
	path string
	mu   *sync.RWMutex
}

type DBStructure struct {
	Tasks map[int]task.Task
}

func (db *DBStructure) GetTasks() map[int]task.Task {
	return db.Tasks
}

var _ Database = (*DB)(nil)

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mu:   &sync.RWMutex{},
	}

	err := db.ensureDB()

	return db, err
}

func (db *DB) ResetDB() error {
	err := os.Remove(db.path)
	if errors.Is(err, ErrNotExist) {
		return nil
	}

	return db.ensureDB()
}

func (db *DB) WriteDB(dbStructure DBStructure) error {
	return db.writeDB(dbStructure)
}

func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	if err := os.WriteFile(db.path, dat, 0644); err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateDB() error {
	return db.createDB()
}

func (db *DB) createDB() error {
	dbStructure := DBStructure{
		Tasks: map[int]Task{},
	}

	return db.writeDB(dbStructure)
}

func (db *DB) EnsureDB() error {
	return db.ensureDB()
}

func (db *DB) ensureDB() error {
	_, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}

	return err
}
func (db *DB) LoadDB() (DBStructure, error) {
	return db.loadDB()
}

func (db *DB) loadDB() (DBStructure, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var dbStructure DBStructure
	dat, err := os.ReadFile(db.path)
	if errors.Is(err, ErrNotExist) {
		return dbStructure, err
	}

	if err := json.Unmarshal(dat, &dbStructure); err != nil {
		return dbStructure, err
	}

	return dbStructure, nil
}
