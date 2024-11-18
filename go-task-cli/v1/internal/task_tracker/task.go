package tasktracker

import "time"

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
	Status      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
