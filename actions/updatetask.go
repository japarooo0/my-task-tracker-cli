package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func Update(id int, name string, f *os.File) {
	defer f.Close()
	var tasks []Task

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if stat.Size() != 0 {
		b, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("file stat: %v", err)
		}

		if err := json.Unmarshal(b, &tasks); err != nil {
			log.Fatalf("file stat: %v", err)
		}
	}

	// check if task does exist
	idx := func() int {
		for i, v := range tasks {
			if id == v.Id {
				return i
			}
		}
		return -1
	}()

	if idx == -1 {
		log.Fatalf("no id found")
	}

	tasks[idx].Name = name
	tasks[idx].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// for i, v := range tasks {
	// 	if v.Id == id {
	// 		tasks[i].Name = name
	// 		tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	// 		fmt.Printf("Update success for %v", v.Id)
	// 	}
	// }

	newJson, err := json.Marshal(&tasks)
	if err != nil {
		log.Fatalf("failed marshal: %v", err)
	}

	if err = os.WriteFile("tasks-data.json", newJson, 0644); err != nil {
		log.Fatalf("failed to write: %v", err)
	}

	fmt.Printf("Update success for id: %d", id)
}
