package state

import (
	"github.com/fakhriaunur/task-cli/internal/database"
)

type State struct {
	db database.Database
}

func NewState(db database.Database) *State {
	return &State{
		db: db,
	}
}
