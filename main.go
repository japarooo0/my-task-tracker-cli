package main

import (
	"log"
	"os"
	"strconv"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
)

// add cli to create task

func main() {

	f, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
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
	case "mark-in-progress":
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("id must be a number: %v", err)
		}
		action.MarkInProgress(v, f)
	case "mark-done":
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("id must be a number: %v", err)
		}
		action.MarkDone(v, f)
	case "list":
		if len(os.Args) == 2 {
			action.ListTasks(f)
		}

		if arg2 := os.Args[2]; arg2 == "done" {
			action.ListDone(arg2, f)
		} else if arg2 == "todo" {
			action.ListTodo(arg2, f)
		} else if arg2 == "in-progress" {
			action.ListInProgress(arg2, f)
		} else {
			log.Fatal("list invalid")
		}
	default:
		log.Fatal("action invalid")
	}

	os.Exit(0)
}
