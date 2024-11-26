package main

import (
	"log"
	"os"

	"github.com/fakhriaunur/task-cli/go-task-cli/v2/adapter/cli"
	"github.com/fakhriaunur/task-cli/go-task-cli/v2/internal/task"
)

func main() {
	cmds := cli.NewCommands()
	cmds.RegisterAllCommands()

	// scmds := cli.NewStandaloneCommands()
	// scmds.RegisterAllStandaloneCommands()

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	taskMapper := task.NewTaskMapper()

	taskDAO := task.NewTaskDAO("./data/task_tracker.json", taskMapper)

	taskRepo, err := task.NewTaskRepo(taskDAO)
	if err != nil {
		log.Fatal(err)
	}

	taskService := task.NewTaskService(taskRepo)

	name := os.Args[1]
	args := os.Args[2:]

	cmd := cli.NewCommand(name, args)

	// scmds.Run(cmd)
	// if err := scmds.Run(cmd); err != nil {
	// 	log.Fatal(err)
	// }
	if err := cmds.Run(taskService, cmd); err != nil {
		log.Fatalf("status: %v\n", err)
	}

}
