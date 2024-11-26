package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/fakhriaunur/task-cli/go-task-cli/v2/internal/task"
)

var (
	ErrTooManyArguments   = errors.New("too many arguments")
	ErrNotEnoughArguments = errors.New("not enough arguments")
)

func handlerAdd(ts task.TaskServicePort, cmd command) error {
	args, err := fetchArgs(cmd, 1)
	if err != nil {
		return err
	}

	desc := args[0]

	newTask, err := ts.AddTask(desc)
	if err != nil {
		return err
	}

	fmt.Println(newTask.ID)

	return nil
}
func handlerUpdate(ts task.TaskServicePort, cmd command) error {
	return nil
}
func handlerDelete(ts task.TaskServicePort, cmd command) error {
	return nil
}
func handlerMarkInProgress(ts task.TaskServicePort, cmd command) error {
	return nil
}
func handlerMarkDone(ts task.TaskServicePort, cmd command) error {
	return nil
}
func handlerList(ts task.TaskServicePort, cmd command) error {
	return nil
}

func handlerHelp(ts task.TaskServicePort, cmd command) error {
	_, err := fetchArgs(cmd, 0)
	if err != nil {
		log.Println("here")
		return err
	}

	fmt.Println("Usage:")

	return nil
}

func fetchArgs(cmd command, n int) ([]string, error) {
	if len(cmd.args) > n {
		return nil, ErrTooManyArguments
	}

	if len(cmd.args) < n {
		return nil, ErrNotEnoughArguments
	}

	return cmd.args, nil
}
