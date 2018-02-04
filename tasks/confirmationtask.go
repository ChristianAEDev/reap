package tasks

import (
	"github.com/manifoldco/promptui"
)

// ConfirmationTask prints a question to the user.
type ConfirmationTask struct {
	RunnableTask
}

func (task ConfirmationTask) GetDescription() (description string) {
	return task.Description
}

func (task ConfirmationTask) Execute() (result Result) {
	question := task.Preferences["Question"].(string)

	answers := []string{"Yes", "No"}

	prompt := promptui.Select{
		Label: question,
		Items: answers,
	}

	_, r, err := prompt.Run()
	if err != nil {
		result.IsSuccessful = false
		result.Message = err.Error()
		return result
	}

	result.IsSuccessful = true
	result.Message = r
	return result
}
