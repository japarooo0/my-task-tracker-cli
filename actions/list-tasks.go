package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func ListTasks(f *os.File) error {
	defer f.Close()
	defer os.Exit(0)
	var tasks []Task

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if err := handlers.ReadAndUnmarshal(f.Name(), &tasks); err != nil {
			return err
		}
	}
	fmt.Print("All Task: \n")
	fmt.Printf("%15s %15s %15s\n", "Name", "Status", "Created At")
	for _, v := range tasks {
		parseTime, _ := time.Parse(time.RFC3339, v.CreatedAt)
		fmt.Printf("%d -- %9s | %s |  %v\n", v.Id, v.Name, v.Status, parseTime.Format("2006-01-02 15:04:05"))
	}

	return nil
}

func ListDone(done string, f *os.File) error {
	defer f.Close()

	var tasks []Task

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if err := handlers.ReadAndUnmarshal(f.Name(), &tasks); err != nil {
			return err
		}
	}

	fmt.Printf("Done task/s: \n")
	var idx int
	for _, v := range tasks {
		if v.Status == done {
			idx++
			fmt.Printf("%d. %9s | Status: %s\n", idx, v.Name, v.Status)
		}
	}

	return nil
}

func ListTodo(todo string, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("file stat: %v ", err)
	}

	if stat.Size() == 0 {
		log.Fatalf("add some task first")
	}

	b, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatalf("read file: %v", err)
	}

	if err := json.Unmarshal(b, &tasks); err != nil {
		log.Fatalf("JSON unmarshal: %v", err)

	}

	var idx int
	fmt.Print("Todo Task \n")
	for _, v := range tasks {
		if v.Status == todo {
			idx++
			fmt.Printf("%d. %9s | Status: %s\n", idx, v.Name, v.Status)
		}
	}

	os.Exit(0)
}

func ListInProgress(progress string, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() == 0 {
		log.Fatalf("add some task first")
	}

	b, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatalf("read file: %v", err)
	}

	if err := json.Unmarshal(b, &tasks); err != nil {
		log.Fatalf("JSON unmarshal: %v", err)
	}

	var idx int
	fmt.Print("In progress task \n")
	for _, v := range tasks {
		if v.Status == progress {
			fmt.Printf("%d. %9s | %s", idx, v.Name, v.Status)
		}
	}

	os.Exit(0)
}
