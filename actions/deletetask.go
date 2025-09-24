package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Delete(id int, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := os.Stat("tasks-data.json")
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() != 0 {
		read, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}

		if err := json.Unmarshal(read, &tasks); err != nil {
			log.Fatalf("JSON unmarshal: %v", err)
		}
	}

	var idx int
	for i, v := range tasks {
		if v.Id == id {
			idx = i
		}
	}

	newTasks := func() []Task {

		if idx == 0 {
			log.Fatalf("no id found")
		}

		tasks[idx] = tasks[len(tasks)-1]
		return tasks[:len(tasks)-1]
	}()

	data, err := json.Marshal(&newTasks)
	if err != nil {
		log.Fatalf("JSON marshal: %v", err)
	}

	if err := os.WriteFile("tasks-data.json", data, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	fmt.Printf("Success delete of id: %v", id)
}
