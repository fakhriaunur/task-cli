package cli

import "github.com/fakhriaunur/task-cli/go-task-cli/v2/internal/task"

type state struct {
	taskService task.TaskServicePort
}

func NewState(ts task.TaskServicePort) *state {
	return &state{
		taskService: ts,
	}
}
