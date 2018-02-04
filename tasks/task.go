package tasks

// Task describes a single step in a plan. This could i.e. be the renaming/deleting of a file.
type Task interface {
	Execute() Result
	GetDescription() string
}

// RunnableTask defines a basic task containing preferences stored as a map.
type RunnableTask struct {
	Description string
	Preferences map[string]interface{}
}

// Result is the outcome of executing a task.
type Result struct {
	Message      string
	IsSuccessful bool
}

// handleError is a convenient function to handle errors occuring during the execution of a task.
// It will set the result to unseccessful and set the error as a message.
func handleError(err error) (result Result) {
	result.IsSuccessful = false
	result.Message = err.Error()
	return result
}
