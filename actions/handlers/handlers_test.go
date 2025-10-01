package handlers_test

import (
	"os"
	"testing"

	action "github.com/japarooo0/my-task-tracker-cli/actions"
	"github.com/japarooo0/my-task-tracker-cli/actions/handlers"
)

func Test_ReadAndUnmarshal(t *testing.T) {
	type expectation struct {
		expectedError string
		expectedData  action.Task
	}
	testCases := []struct {
		desc        string
		data        string
		fileName    string
		expectation expectation
	}{
		{
			desc:        "invalid json",
			data:        ``,
			fileName:    "test.json",
			expectation: expectation{expectedError: "JSON unmarshal: unexpected end of JSON input"},
		},

		{
			desc:     "success",
			fileName: "test.json",
			data: `[{
				"id": 10,
				"name": "Sample task",
				"status": "todo",
				"created_at": "2006-01-02 15:04:05",
				"updated_at": "2006-01-02 15:04:05"
			}]`,
			expectation: expectation{expectedData: action.Task{
				Id:        10,
				Name:      "Sample task",
				Status:    "todo",
				CreatedAt: "2006-01-02 15:04:05",
				UpdatedAt: "2006-01-02 15:04:05",
			}},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var tasks []action.Task
			f, _ := handlers.CreateTestFile(tC.fileName, tC.data)

			err := handlers.ReadAndUnmarshal(f.Name(), &tasks)

			if err != nil && err.Error() != tC.expectation.expectedError {
				t.Errorf("expected: %v, got: %v", tC.expectation.expectedError, err)
			}

			if tasks != nil && tasks[0] != tC.expectation.expectedData {
				t.Errorf("expected: %v, got: %v", tC.expectation.expectedData, tasks[0])
			}

			os.Remove(f.Name())
		})
	}
}

func Test_MarshalAndWrite(t *testing.T) {
	testCases := []struct {
		desc     string
		data     []action.Task
		expected string
	}{
		{
			desc: "success",
			data: []action.Task{{
				Id:        10,
				Name:      "Sample task",
				Status:    "todo",
				CreatedAt: "2006-01-02 15:04:05",
				UpdatedAt: "2006-01-02 15:04:05",
			}},
		},
		{
			desc: "invalid to marshal",
			data: []action.Task{{
				Name:      "Sample task",
				Status:    "todo",
				CreatedAt: "2006-01-02 15:04:05",
				UpdatedAt: "2006-01-02 15:04:05",
			}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			f, _ := handlers.CreateTestFile("test.json", `[]`)
			err := handlers.MarshalAndWrite(f.Name(), &tC.data)

			if err != nil {
				t.Errorf("expected: %v, got: %v", nil, err)
			}

			os.Remove(f.Name())
		})
	}
}
