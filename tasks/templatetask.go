package tasks

import (
	"bufio"
	"fmt"
	"os"
)

type TemplateTask struct {
	RunnableTask
}

func (task TemplateTask) GetDescription() (description string) {
	return task.Description
}

func (task TemplateTask) Execute() (result Result) {
	// Path where the template will be saved
	filePath := task.Preferences["FilePath"].(string)
	// Template stores an array. Each item in the array is one line.
	template := task.Preferences["Template"].([]interface{})

	// Write the template file to the disk
	file, err := os.Create(filePath)
	if err != nil {
		return handleError(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range template {
		fmt.Fprintln(w, line)
	}
	w.Flush()

	result.IsSuccessful = true
	result.Message = "Template expanded"
	return result
}
