package main

import "github.com/fakhriaunur/task-cli/internal/database"

type stateManager struct {
	db *database.DB
}

func NewState(db *database.DB) *stateManager {
	return &stateManager{
		db: db,
	}
}
