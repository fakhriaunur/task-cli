package main

import (
	"log"
	"os"

	"github.com/fakhriaunur/task-cli/cli"
	"github.com/fakhriaunur/task-cli/internal/persistence"
	tt "github.com/fakhriaunur/task-cli/internal/task_tracker"
)

func main() {
	db, err := persistence.NewJSONRepo("./data/task_tracker.json")
	if err != nil {
		log.Fatal()
	}

	taskService := tt.NewTaskService(db)
	cmds := cli.NewCommands()

	cmds.Register("add", cli.HandlerAdd)
	cmds.Register("update", cli.HandlerUpdate)
	cmds.Register("delete", cli.HandlerDelete)
	cmds.Register("mark-in-progress", cli.HandlerMarkInProgress)
	cmds.Register("mark-done", cli.HandlerMarkDone)
	cmds.Register("list", cli.HandlerList)
	cmds.Register("help", cli.HandlerHelp)
	cmds.Register("reset", cli.HandlerReset)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(taskService, cmd); err != nil {
		return
	}
}
