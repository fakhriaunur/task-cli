package main

import (
	"log"
	"os"
)

func main() {
	cmds := NewCommands()
	cmds.register("add", handlerAdd)
	cmds.register("update", handlerUpdate)
	cmds.register("delete", handlerDelete)
	cmds.register("mark-in-progress", handlerMarkInProgress)
	cmds.register("mark-done", handlerMarkDone)
	cmds.register("list", handlerList)
	cmds.register("help", handlerHelp)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err := cmds.run(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
