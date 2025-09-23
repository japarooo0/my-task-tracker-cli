package main

import (
	"fmt"
	"log"
	"os"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
)

// add cli to create task

func main() {

	f, err := os.OpenFile("tasks-data.json", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("creating file: %w", err))
	}

	switch os.Args[1] {
	case "add":
		action.Add(os.Args[2], f)
	default:
		log.Fatal("action invalid")
	}

	defer os.Exit(0)
}
