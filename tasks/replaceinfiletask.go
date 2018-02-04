package tasks

import (
	"bytes"
	"io/ioutil"
)

// ReplaceInFileTask takes a text file an replaces all occurences of a given string with another
// given string.
type ReplaceInFileTask struct {
	RunnableTask
}

func (task ReplaceInFileTask) GetDescription() (description string) {
	return task.Description
}

func (task ReplaceInFileTask) Execute() (result Result) {
	filePath := task.Preferences["FilePath"].(string)
	replace := task.Preferences["Replace"].(string)
	with := task.Preferences["With"].(string)

	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return handleError(err)
	}

	output := bytes.Replace(input, []byte(replace), []byte(with), -1)

	if err = ioutil.WriteFile(filePath, output, 0666); err != nil {
		return handleError(err)
	}

	result.IsSuccessful = true
	return result
}
