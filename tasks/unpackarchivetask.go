package tasks

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// UnpackArchiveTask extracts the files from a given zip archive to a given folder.
type UnpackArchiveTask struct {
	RunnableTask
}

func (task UnpackArchiveTask) GetDescription() (description string) {
	return task.Description
}

func (task UnpackArchiveTask) Execute() (result Result) {
	src := task.Preferences["FilePath"].(string)
	dest := task.Preferences["DestinationPath"].(string)

	// Create a reader for the specified file
	reader, err := zip.OpenReader(src)
	if err != nil {
		result.IsSuccessful = false
		result.Message = err.Error()
		return result
	}
	defer reader.Close()

	// Range over the content of the archive
	for _, f := range reader.File {

		rc, err := f.Open()
		if err != nil {
			result.IsSuccessful = false
			result.Message = err.Error()
			return result
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		// Check if the current file points to a file or a folder
		if f.FileInfo().IsDir() {
			// Create the folder
			os.MkdirAll(path, os.ModePerm)
		} else {
			// Create the file
			var fileDir string
			if lastIndex := strings.LastIndex(path, string(os.PathSeparator)); lastIndex > -1 {
				fileDir = path[:lastIndex]
			}

			err = os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				result.IsSuccessful = false
				result.Message = err.Error()
				return result
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				result.IsSuccessful = false
				result.Message = err.Error()
				return result
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				result.IsSuccessful = false
				result.Message = err.Error()
				return result
			}
		}
	}

	result.IsSuccessful = true
	result.Message = "Unpacked"
	return result
}
