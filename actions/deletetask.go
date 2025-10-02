package action

import (
	"fmt"
	"os"

	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Delete(id int, f *os.File) error {
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
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	if err := handlers.MarshalAndWrite(f.Name(), &tasks); err != nil {
		return err
	}

	fmt.Printf("Success delete of id: %v", id)
	return nil
}
