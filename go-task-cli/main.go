package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fakhriaunur/task-cli/internal/repl"
)

func main() {
	fmt.Println("Welcome to the Task CLI")
	fmt.Println("=======================")

	err := repl.Start(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatalf("Error starting the REPL: %v", err)
	}
}
