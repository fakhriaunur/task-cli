package cli

import (
	"errors"

	"github.com/fakhriaunur/task-cli/go-task-cli/v2/internal/task"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCmds map[string]func(task.TaskServicePort, command) error
}

func NewCommands() *commands {
	return &commands{
		registeredCmds: make(map[string]func(task.TaskServicePort, command) error),
	}
}

func (c *commands) register(
	name string,
	f func(task.TaskServicePort, command) error,
) {
	c.registeredCmds[name] = f
}

func (c *commands) run(
	ts task.TaskServicePort,
	cmd command,
) error {
	f, ok := c.registeredCmds[cmd.name]
	if !ok {
		return errors.New("command is not registered")
	}

	return f(ts, cmd)
}
