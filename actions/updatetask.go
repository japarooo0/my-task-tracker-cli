package action

import (
	"fmt"
	"os"
	"time"

	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Update(id int, name string, f *os.File) error {
	defer f.Close()
	var tasks []Task

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if err := handlers.ReadAndUnmarshal(f.Name(), &tasks); err != nil {
			return err
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
		return fmt.Errorf("no id found")
	}

	if err := tasks[idx].Validate(); err != nil {
		return err
	}
	tasks[idx].Name = name
	tasks[idx].UpdatedAt = time.Now().Format(time.RFC3339)

	// for i, v := range tasks {
	// 	if v.Id == id {
	// 		tasks[i].Name = name
	// 		tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	// 		fmt.Printf("Update success for %v", v.Id)
	// 	}
	// }

	if err := handlers.MarshalAndWrite(f.Name(), &tasks); err != nil {
		return err
	}

	fmt.Printf("Update success for id: %d", id)
	return nil
}
