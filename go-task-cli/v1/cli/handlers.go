package cli

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

const (
	ArgumentNotEnough = "not enough argument"
	ArgumentTooMany   = "too many arguments"
)

// validate here
func HandlerAdd(ts tt.TaskServiceCommonPort, cmd Command) error {
	args, err := fetchArgs(cmd, 1)
	if err != nil {
		log.Println(err)
		return err
	}

	desc := args[0]

	task, err := ts.Add(desc)
	if err != nil {
		fmt.Println("err")
		return err
	}

	fmt.Printf("Successfully add a task with ID=%d\n", task.ID)

	return nil
}
func HandlerUpdate(ts tt.TaskServiceCommonPort, cmd Command) error {
	args, err := fetchArgs(cmd, 2)
	if err != nil {
		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	desc := args[1]

	if err := ts.Update(id, desc); err != nil {
		return err
	}

	fmt.Printf("Successfully updated a task with ID=%d", id)

	return nil
}
func HandlerDelete(ts tt.TaskServiceCommonPort, cmd Command) error {
	args, err := fetchArgs(cmd, 1)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Successfully deleted a task with ID=%d", id)

	return nil
}

func HandlerMarkInProgress(ts tt.TaskServiceCommonPort, cmd Command) error {
	args, err := fetchArgs(cmd, 1)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Successfully marked in-progress a task with ID=%d", id)

	return nil
}

func HandlerMarkDone(ts tt.TaskServiceCommonPort, cmd Command) error {
	args, err := fetchArgs(cmd, 1)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Successfully marked done a task with ID=%d", id)

	return nil

}
func HandlerList(ts tt.TaskServiceCommonPort, cmd Command) error {
	var tasks []tt.Task
	var err error

	switch len(cmd.Args) {
	case 0:
		tasks, err = ts.ListAll()
	case 1:
		args, err := fetchArgs(cmd, 1)
		if err != nil {
			return err
		}

		status := args[0]

		switch status {
		case tt.StatusDone, tt.StatusInProgress, tt.StatusTodo:
			tasks, err = ts.ListByStatus(status)
			if err == nil && len(tasks) == 0 {
				fmt.Printf("No %s tasks currently", status)
			}
		}
	}

	if err != nil {
		return err
	}

	sortLists(tasks)
	printLists(tasks)

	return nil
}

func sortLists(tasks []tt.Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
}

func printLists(tasks []tt.Task) {
	for _, task := range tasks {
		fmt.Printf("%d\n", task.ID)
		fmt.Printf("%s\n", task.Status)
		fmt.Printf("%s\n", task.Description)
		fmt.Printf("%v\n", task.CreatedAt)
		fmt.Printf("%v\n", task.UpdatedAt)
	}
}

func HandlerHelp(ts tt.TaskServiceCommonPort, cmd Command) error {
	_, err := fetchArgs(cmd, 0)
	if err != nil {
		return err
	}

	fmt.Println(ts.Help())
	printHelp()

	return nil
}

func HandlerReset(ts tt.TaskServiceCommonPort, cmd Command) error {
	_, err := fetchArgs(cmd, 0)
	if err != nil {
		return err
	}

	ts.Reset()

	fmt.Println("Successfully resetted the DB")

	return nil
}

func printHelp() {
	fmt.Println("Usage: ")
	fmt.Println("add")
	fmt.Println("update <id> <desc>")
	fmt.Println("delete <id>")
	fmt.Println("mark-in-progress <id>")
	fmt.Println("mark-done <id>")
	fmt.Println("list")
	fmt.Println("list <status>")
	fmt.Println("help")
}

func fetchArgs(cmd Command, n int) ([]string, error) {
	if len(cmd.Args) < n {
		return nil, errors.New(ArgumentNotEnough)
	}
	if len(cmd.Args) > n {
		return nil, errors.New(ArgumentTooMany)
	}

	return cmd.Args, nil
}
