package action

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (t *Task) Validate() (err error) {
	if t.Name == "" {
		err = errors.Join(err, fmt.Errorf("name is empty"))
	}
	if t.Status == "" {
		err = errors.Join(err, fmt.Errorf("status is empty"))
	}

	_, timeErr := time.Parse(time.RFC3339, t.CreatedAt)
	if timeErr != nil {
		err = errors.Join(err, timeErr)
	}

	_, updatedTimeErr := time.Parse(time.RFC3339, t.UpdatedAt)
	if updatedTimeErr != nil {
		err = errors.Join(err, updatedTimeErr)
	}

	return err
}

type Actions interface {
	Add(string)
	Update(int, string)
	Delete(int)
	MarkProgress(int)
	MarkDone(int)
	List()
}
