package cli

import "errors"

// TODO: Make decorator / middleware. Start from the base standaloneCommands, then
// add states/ services as necessary
type Commands interface {
	register(string, func(command) error)
	Run(command) error
}

type standaloneCommands struct {
	registeredCmds map[string]func(command) error
}

// Register implements Commands.
func (s *standaloneCommands) register(
	name string,
	f func(command) error,
) {
	s.registeredCmds[name] = f
}

// Run implements Commands.
func (s *standaloneCommands) Run(cmd command) error {
	f, ok := s.registeredCmds[cmd.name]
	if !ok {
		return errors.New("couldn't find command")
	}

	return f(cmd)
}

func NewStandaloneCommands() *standaloneCommands {
	return &standaloneCommands{
		registeredCmds: make(map[string]func(command) error),
	}
}
