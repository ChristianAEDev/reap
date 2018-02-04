package tasks

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

// ExecCommandTask executes a specified console command
type ExecCommandTask struct {
	RunnableTask
}

func (task ExecCommandTask) GetDescription() (description string) {
	return task.Description
}

func (t ExecCommandTask) Execute() (result Result) {
	command := t.Preferences["Command"].(string)
	var args []string
	for _, value := range t.Preferences["Args"].([]interface{}) {
		args = append(args, value.(string))
	}
	// Optional dir parameter to define in which folder to execute the command
	dir, ok := t.Preferences["Dir"].(string)
	if ok {
		dir = t.Preferences["Dir"].(string)
	} else {
		dir = ""
	}

	// Create the command
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	var stdBuffer bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		return Result{
			err.Error(),
			false,
		}
	}

	// Output of the comman executed. Could later be used to have some logic running on it to
	// determine if the execution was successful.
	//output := stdBuffer.String()

	result.IsSuccessful = true
	result.Message = "The command \"" + command + "\"" + " was executed successfully."

	return result
}
