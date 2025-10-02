package action_test

import (
	"os"
	"testing"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Test_DeleteTask(t *testing.T) {
	testCases := []struct {
		desc     string
		jsonData string
		id       int
		expected string
	}{
		{
			desc: "invalid id",
			jsonData: `[{
				"id": 1,
				"name": "task to delete",
				"status": "todo",
				"created_at": "2025-10-01T13:51:16+08:00",
				"updated_at": "2025-10-01T13:51:16+08:00"
			}]`,
			id:       0,
			expected: "no id found",
		},
		{
			desc: "success",
			jsonData: `[{
				"id": 1,
				"name": "task to delete",
				"status": "todo",
				"created_at": "2025-10-01T13:51:16+08:00",
				"updated_at": "2025-10-01T13:51:16+08:00"
			}]`,
			id: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f, err := handlers.CreateTestFile("delete-test.json", tC.jsonData)
			if err != nil {
				t.Errorf("failed to create test file")
			}
			err = action.Delete(tC.id, f)
			defer os.Remove(f.Name())

			if err != nil && err.Error() != tC.expected {
				t.Errorf("expected: %v, got: %v", tC.expected, err.Error())
			}

		})
	}
}
