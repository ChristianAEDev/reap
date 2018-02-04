package tasks

// ServiceTask allows to start/stop a service installed on the server this tools in run on.
type ServiceTask struct {
	RunnableTask
}

func (task ServiceTask) GetDescription() (description string) {
	return task.Description
}

// Execute renames a file identified by its full path and (old) name.
func (task ServiceTask) Execute() (result Result) {
	result.IsSuccessful = false
	result.Message = "This task is not implemented on this platform."
	return result
}
