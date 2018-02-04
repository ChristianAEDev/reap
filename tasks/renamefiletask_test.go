package tasks

import (
	"os"
	"testing"
)

const pathTestData = "testing"

func TestRenameFile(t *testing.T) {
	methodName := "TestRenameFile"
	path := pathTestData + string(os.PathSeparator) + methodName
	oldName := "OldName.txt"
	newName := "NewName.txt"

	fullPathOldName := path + string(os.PathSeparator) + oldName
	fullPathNewName := path + string(os.PathSeparator) + newName

	// Create test directory
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	// Create file to rename
	file, err := os.Create(fullPathOldName)
	if err != nil {
		t.Error(err)
	}
	file.Close()

	// Create the RenameFileTask
	task := RenameFileTask{
		RunnableTask{
			Preferences: map[string]interface{}{
				"FilePath": path,
				"OldName":  oldName,
				"NewName":  newName,
			},
		},
	}
	// Execute the RenameFileTask
	result := task.Execute()
	// Check the result
	if !result.IsSuccessful {
		t.Errorf(result.Message)
	} else {
		if _, err := os.Stat(fullPathNewName); os.IsNotExist(err) {
			// Test failed. The expected file does not exist
			t.Error(err)
		}
	}
	// Clean up
	os.RemoveAll(path)
	os.Remove(pathTestData)
}
