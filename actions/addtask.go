package action

import (
	"fmt"
	"os"
	"time"

	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Add(name string, f *os.File) error {

	var tasks []Task
	defer f.Close()

	isEmpty := handlers.IsFileEmpty(f.Name())

	if !isEmpty {
		if errs := handlers.ReadAndUnmarshal(f.Name(), &tasks); errs != nil {
			return errs
		}
	}

	t := time.Now().Format(time.RFC3339)

	task := Task{Name: name, Status: "todo", CreatedAt: t, UpdatedAt: t}
	if len(tasks) != 0 {
		task.Id = tasks[len(tasks)-1].Id + 1
	} else {
		task.Id = 1
	}

	tasks = append(tasks, task)

	for _, v := range tasks {
		errs := v.Validate()
		if errs != nil {
			return errs
		}
	}

	if errs := handlers.MarshalAndWrite(f.Name(), &tasks); errs != nil {
		return errs
	}

	fmt.Print("Task added")
	return nil
}
