package action_test

import (
	"os"
	"testing"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
)

func Test_AddTask(t *testing.T) {
	testCases := []struct {
		desc     string
		taskName string
		expected string
	}{
		{
			desc:     "invalid name",
			taskName: "",
			expected: "name is empty",
		},
		{
			desc:     "success",
			taskName: "Paragliding tomorrow",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f, errs := os.OpenFile("add-test.json", os.O_CREATE|os.O_RDWR, 0644)
			if errs != nil {
				t.Errorf("error file creation: %v", errs)
			}

			err := action.Add(tC.taskName, f)

			if err != nil && err.Error() != tC.expected {
				t.Errorf("expected: %v, got: %v", tC.expected, err)
			}

			os.Remove(f.Name())
		})
	}
}
