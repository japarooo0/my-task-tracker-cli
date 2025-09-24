package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
)

// add cli to create task

func main() {

	f, err := os.OpenFile("tasks-data.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Errorf("creating file: %w", err))
	}

	switch os.Args[1] {
	case "add":
		action.Add(os.Args[2], f)
	case "update":
		v, err := strconv.Atoi(os.Args[2])
		task := os.Args[3]
		if err != nil {
			log.Fatalf("id must be a number: %v", err)
		}
		action.Update(v, task, f)
	case "delete":
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("id must be a number: %v", err)
		}
		action.Delete(v, f)
	default:
		log.Fatal("action invalid")
	}

	defer os.Exit(0)
}
