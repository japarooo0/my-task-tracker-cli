package action

import (
	"fmt"
	"os"
	"time"

	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func MarkInProgress(id int, f *os.File) error {
	defer f.Close()
	var tasks []Task

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if err := handlers.ReadAndUnmarshal(f.Name(), &tasks); err != nil {
			return err
		}
	}

	for i, v := range tasks {
		if v.Id == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)

		}
	}

	if err := handlers.MarshalAndWrite(f.Name(), &tasks); err != nil {
		return err
	}

	fmt.Println("Task marking success")
	return nil
}
func MarkDone(id int, f *os.File) error {
	defer f.Close()
	var tasks []Task

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if err := handlers.ReadAndUnmarshal(f.Name(), &tasks); err != nil {
			return err
		}
	}

	for i, v := range tasks {
		if v.Id == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
		}
	}

	if err := handlers.MarshalAndWrite(f.Name(), &tasks); err != nil {
		return err
	}

	fmt.Println("Task marking success")
	return nil
}
