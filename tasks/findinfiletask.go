package tasks

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// FindInFileTask searches in the given file for a given string.
type FindInFileTask struct {
	RunnableTask
}

func (task FindInFileTask) GetDescription() (description string) {
	return task.Description
}

func (task FindInFileTask) Execute() (result Result) {
	filePath := task.Preferences["FilePath"].(string)
	query := task.Preferences["Query"].(string)

	file, err := os.Open(filePath)
	if err != nil {
		return handleError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hits := 0
	ln := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, query) {
			hits++
		}
		ln++
	}

	result.IsSuccessful = true
	result.Message = "Found " + strconv.Itoa(hits) + " hit(s)"
	return result
}
