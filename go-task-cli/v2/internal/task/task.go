package task

import "time"

const (
	statusTodo       = "todo"
	statusInProgress = "in-progress"
	statusDone       = "done"
)

var (
	timeNow  = time.Now()
	timeZero = time.Time{}
)

type task struct {
	ID          int
	Status      string
	Description string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
