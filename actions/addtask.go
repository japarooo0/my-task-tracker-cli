package action

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

func Add(name string, f *os.File) {
	var tasks []Task
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}

	if fileInfo.Size() != 0 {
		fileValue, err := os.ReadFile("tasks-data.json")
		if err != nil {
			log.Fatalf("read file: %v", err)
		}
		err = json.Unmarshal(fileValue, &tasks)
		if err != nil {
			log.Fatalf("unmarshal fail: %v", err)
		}
	}

	r := rand.New(rand.NewSource(99))
	t := time.Now()
	if err != nil {
		log.Fatalf("parse time fail: %v", err)
	}

	task := Task{Name: name, Status: "todo", CreatedAt: t.Format("2006-01-02 15:04:05"), UpdatedAt: t.Format("2006-01-02 15:04:05")}
	task.Id = r.Intn(999)

	for _, v := range tasks {
		if v.Id == task.Id {
			task.Id = r.Intn(999)
			continue
		}
	}

	tasks = append(tasks, task)

	d, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("marshal fail: %v", err)
	}

	if _, err := f.Write([]byte(d)); err != nil {
		log.Fatalf("failed to write: %v", err)
	}

	f.Close()
}
