package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCmds map[string]func(*stateManager, command) error
}

func NewCommands(s *stateManager) *commands {
	handlers := make(map[string]func(*stateManager, command) error)

	return &commands{
		registeredCmds: handlers,
	}
}

func (c *commands) register(name string, handler func(*stateManager, command) error) {
	c.registeredCmds[name] = handler
}

func (c *commands) run(s *stateManager, cmd command) error {
	handler, ok := c.registeredCmds[cmd.name]
	if !ok {
		return errors.New("handler is not exist")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
