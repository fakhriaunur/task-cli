package task

import (
	"encoding/json"
	"os"
)

const filename = "tasks.json"

func SaveTasks(tasks []Task) error {
	file, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func LoadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.OpenFile(
		filename,
		os.O_RDONLY|os.O_CREATE,
		0755,
	)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		return tasks, err
	}

	return tasks, nil
}
