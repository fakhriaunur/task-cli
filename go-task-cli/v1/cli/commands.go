package cli

import (
	"errors"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

type Command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCmds map[string]func(tt.TaskServiceCommonPort, Command) error
}

func NewCommands() *commands {
	handler := make(map[string]func(tt.TaskServiceCommonPort, Command) error)

	return &commands{
		registeredCmds: handler,
	}
}

func (c *commands) Register(name string, handler func(tt.TaskServiceCommonPort, Command) error) {
	c.registeredCmds[name] = handler
}

func (c *commands) Run(ts tt.TaskServiceCommonPort, cmd Command) error {
	handler, ok := c.registeredCmds[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return handler(ts, cmd)
	// if err := handler(ts, cmd); err != nil {
	// 	return err
	// }

	// return nil
}
