package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func IsFileEmpty(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalf("file stat: %v", err)
	}
	return stat.Size() == 0
}

func ReadAndUnmarshal[T any](path string, tasks *[]T) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %v", err)
	}

	if err := json.Unmarshal(b, &tasks); err != nil {
		return fmt.Errorf("JSON unmarshal: %v", err)
	}

	return nil
}

func MarshalAndWrite[T any](path string, tasks *[]T) error {
	b, err := json.Marshal(&tasks)
	if err != nil {
		return fmt.Errorf("JSON marshal: %v", err)
	}

	if err := os.WriteFile(path, b, 0644); err != nil {
		return fmt.Errorf("write file: %v", err)
	}

	return nil
}

func CreateTestFile(path string, data string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	f.Write([]byte(data))
	defer f.Close()
	return f, nil
}
