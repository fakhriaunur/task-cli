package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCmds map[string]func(command) error
}

func NewCommands() *commands {
	handlers := make(map[string]func(command) error)

	return &commands{
		registeredCmds: handlers,
	}
}

func (c *commands) register(name string, f func(command) error) {
	c.registeredCmds[name] = f
}

func (c *commands) run(cmd command) error {
	handler, ok := c.registeredCmds[cmd.name]
	if !ok {
		return errors.New("unknown command")
	}

	err := handler(cmd)
	if err != nil {
		return err
	}

	return nil
}
