package main

import (
	"log"
	"os"

	"github.com/fakhriaunur/task-cli/internal/database"
	"github.com/fakhriaunur/task-cli/internal/state"
	"github.com/fakhriaunur/task-cli/internal/task"
)

func main() {
	db, err := database.NewDB("task-cli.json")
	if err != nil {
		log.Fatal(err)
	}

	tm := task.NewTaskManager(db)

	programState := state.NewState(db, tm)

	cmds := NewCommands(programState)

	cmds.register("add", handlerAdd)
	cmds.register("update", handlerUpdate)
	cmds.register("delete", handlerDelete)
	cmds.register("mark-in-progress", handlerMarkInProgress)
	cmds.register("mark-done", handlerMarkDone)
	cmds.register("list", handlerList)
	cmds.register("help", handlerHelp)

	if len(os.Args) < 2 {
		log.Fatal("error: not enough argument")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
