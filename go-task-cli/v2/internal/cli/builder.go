package cli

func middlewareWithState(handler func(s *state, cmd command) error) func(command) error {
	return func(cmd command) error {
		return handler(s, cmd)
	}
}
