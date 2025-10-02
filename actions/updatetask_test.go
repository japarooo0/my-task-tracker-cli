package action_test

import (
	"os"
	"testing"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		jsonData string
		id       int
		taskName string
		expected string
	}{
		{
			desc:     "invalid json",
			jsonData: `{`,
			id:       1,
			taskName: "A sample task",
			expected: "JSON unmarshal: unexpected end of JSON input",
		},
		{
			desc: "invalid id",
			jsonData: `[{
				"id": 1,
				"name": "Initial task",
				"status": "todo",
				"created_at": "2025-10-01T13:51:16+08:00",
				"updated_at": "2025-10-01T13:51:16+08:00"
			}]`,
			id:       0,
			taskName: "A sample task",
			expected: "no id found",
		},
		{
			desc:     "invalid name",
			id:       1,
			taskName: "",
			jsonData: `[{
				"id": 1,
				"name": "Initial task",
				"status": "todo",
				"created_at": "2025-10-01T13:51:16+08:00",
				"updated_at": "2025-10-01T13:51:16+08:00"
			}]`,
			expected: "name is empty",
		},
		{
			desc:     "success",
			id:       1,
			taskName: "A sample task",
			jsonData: `[{
				"id": 1,
				"name": "Initial task",
				"status": "todo",
				"created_at": "2025-10-01T13:51:16+08:00",
				"updated_at": "2025-10-01T13:51:16+08:00"
			}]`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f, errs := handlers.CreateTestFile("update-test.json", tC.jsonData)
			if errs != nil {
				t.Errorf("file creation: %v", errs)
			}
			err := action.Update(tC.id, tC.taskName, f)
			defer os.Remove(f.Name())

			if err != nil && err.Error() != tC.expected {
				t.Errorf("expected: %v, got: %v", tC.expected, err)
			}

		})
	}
}
