package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fakhriaunur/task-cli/internal/task"
)

// TODO: Implement handlers: add, update, delete, marking-in, list-all, list-by-status
func handlerAdd(cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an argument")
	}

	description := cmd.args[0]
	_, err := task.Add(description)
	if err != nil {
		return err
	}

	return nil
}

func handlerUpdate(cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("expecting <id> <description>")
	}

	id, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	desc := cmd.args[1]

	return task.Update(id, desc)
}

func handlerDelete(cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an <id>")
	}

	id, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	return task.Delete(id)
}

func handlerMarkInProgress(cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an <id>")
	}

	id, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	return task.MarkInProgress(id)
}

func handlerMarkDone(cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an <id>")
	}

	id, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	return task.MarkDone(id)
}

func handlerList(cmd command) error {
	switch len(cmd.args) {
	case 0:
		tasks, err := task.ListAll()
		if err != nil {
			return err
		}
		for _, task := range tasks {
			printTask(task)
		}
	case 1:
		status := cmd.args[0]
		tasks, err := task.ListByStatus(status)
		if err != nil {
			return err
		}
		for _, task := range tasks {
			printTask(task)
		}
	default:
		return errors.New("expecing 0 or 1 argument")
	}

	return nil
}

func printTask(task task.Task) {
	fmt.Println(task.ID)
	fmt.Println(task.Description)
	fmt.Println(task.Status)
	fmt.Println(task.CreatedAt)
	fmt.Println(task.UpdatedAt)

}

func handlerHelp(cmd command) error {
	if len(cmd.args) > 0 {
		fmt.Println("don't need any argument")
	}

	fmt.Println("Usage: ")
	fmt.Println("add")
	fmt.Println("update")
	fmt.Println("delete")
	fmt.Println("mark-done")
	fmt.Println("mark-in-progress")
	fmt.Println("list")
	fmt.Println("list <status>")

	return nil
}
