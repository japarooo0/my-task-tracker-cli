package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func MarkInProgress(id int, f *os.File) {
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

	for i, v := range tasks {
		if v.Id == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

		}
	}

	data, err := json.Marshal(&tasks)
	if err != nil {
		log.Fatalf("JSON marshal: %v", err)
	}

	if err := os.WriteFile("tasks-data.json", data, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	fmt.Println("Task marking success")
}
func MarkDone(id int, f *os.File) {
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

	for i, v := range tasks {
		if v.Id == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		}
	}

	data, err := json.Marshal(&tasks)
	if err != nil {
		log.Fatalf("JSON marshal: %v", err)
	}

	if err := os.WriteFile("tasks-data.json", data, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	fmt.Println("Task marking success")
}
