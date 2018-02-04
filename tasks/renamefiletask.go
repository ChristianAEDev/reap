package tasks

import (
	"os"
)

// RenameFileTask renames a file.
type RenameFileTask struct {
	RunnableTask
}

func (task RenameFileTask) GetDescription() (description string) {
	return task.Description
}

// Execute renames a file identified by its full path and (old) name.
func (task RenameFileTask) Execute() (result Result) {
	// Extract all needed preferences
	filePath := task.Preferences["FilePath"].(string)
	oldName := task.Preferences["OldName"].(string)
	newName := task.Preferences["NewName"].(string)
	// Build the full paths from the config
	fullPathOld := filePath + string(os.PathSeparator) + oldName
	fullPathNew := filePath + string(os.PathSeparator) + newName

	// Make sure the file exists
	if _, err := os.Stat(fullPathOld); os.IsNotExist(err) {
		result.Message = "The path \"" + fullPathOld + "\" does not exist"
		result.IsSuccessful = false
		// The file to rename does not exists
		return result
	}

	err := os.Rename(fullPathOld, fullPathNew)
	result.IsSuccessful = err == nil
	if err != nil {
		result.Message = err.Error()
	}
	result.Message = "Renamed file \"" + fullPathOld + "\" to \"" + fullPathNew + "\""

	return result
}
