package tasks

import (
	"os"
)

// DeleteTask allows to delete a given file/folder.
type DeleteTask struct {
	RunnableTask
}

func (task DeleteTask) GetDescription() (description string) {
	return task.Description
}

func (task DeleteTask) Execute() (result Result) {
	path := task.Preferences["Path"].(string)

	f, err := os.Stat(path)
	if err != nil {
		result.IsSuccessful = false
		result.Message = err.Error()
		return result
	}

	if f.IsDir() {
		err := os.RemoveAll(path)
		if err != nil {
			return handleError(err)
		}
	} else {
		err := os.Remove(path)
		if err != nil {
			return handleError(err)
		}
	}

	result.IsSuccessful = true
	result.Message = path + " deleted"
	return result
}
