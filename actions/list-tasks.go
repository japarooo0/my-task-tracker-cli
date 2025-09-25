package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListTasks(f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := os.Stat("tasks-data.json")
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() != 0 {
		b, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}

		if err := json.Unmarshal(b, &tasks); err != nil {
			log.Fatalf("JSON unmarshal: %v", err)
		}
	}

	fmt.Print("All Task: \n")
	for _, v := range tasks {
		fmt.Printf("%3d. %9s | Status: %s\n", v.Id, v.Name, v.Status)
	}

	os.Exit(0)
}

func ListDone(done string, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := os.Stat("tasks-data.json")
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() != 0 {
		b, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}

		if err := json.Unmarshal(b, &tasks); err != nil {
			log.Fatalf("JSON unmarshal: %v", err)
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

}

func ListTodo(todo string, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := os.Stat("tasks-data.json")
	if err != nil {
		log.Fatalf("file stat: %v ", err)
	}

	if stat.Size() != 0 {
		b, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}

		if err := json.Unmarshal(b, &tasks); err != nil {
			log.Fatalf("JSON unmarshal: %v", err)

		}
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

	stat, err := os.Stat("tasks-data.json")
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() != 0 {
		b, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}

		if err := json.Unmarshal(b, &tasks); err != nil {
			log.Fatalf("JSON unmarshal: %v", err)
		}
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
