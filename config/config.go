// Package config holds functions responsible to handle a config file. The main functionallity
// is to create the correct tasks.
package config

import (
	"encoding/json"
	"log"

	"github.com/ChristianAEDev/reap/tasks"
)

// AppConfig contains the config created from the json file
var AppConfig Config

// Config holds the instructions on what actions to perform. Each plan consists out of one ore more
// tasks (steps) to execute
type Config struct {
	Plans []Plan
}

// Plan holds one or more tasks. The tasks in a plan are executed according to their order in the
// config file.
type Plan struct {
	Name  string
	Tasks []tasks.Task
}

// ExtractConfig takes a byte array that is formatted as a JSON and extracts all the plans it defines
func ExtractConfig(configJSON []byte) (config Config) {
	// Read the plans from the config
	var configMap map[string]interface{}

	// Parse the config file
	if err := json.Unmarshal([]byte(configJSON), &configMap); err != nil {
		log.Panic(err)
	}
	return ParseConfig(configMap)
}

// ParseConfig iterates of the the content of the config JSON file and extracts it's plans.
func ParseConfig(configMap map[string]interface{}) (config Config) {
	// Range over first layer of the JSON file
	for _, value := range configMap {
		// Range over each plan in the JSON file
		config.Plans = parsePlans(value.([]interface{}))
	}
	return config
}

// ParsePlans extracts the plas from the given plans map (JSON representation).
func parsePlans(plansSlice []interface{}) (plans []Plan) {
	for _, value := range plansSlice {
		plan := parsePlan(value.(map[string]interface{}))
		plans = append(plans, plan)
	}
	return plans
}

// ParsePlan extracts a single plan from the given plan map (JSON representation).
func parsePlan(planMap map[string]interface{}) Plan {
	plan := Plan{}

	for key, value := range planMap {
		switch key {
		case "Name":
			plan.Name = value.(string)
		case "Tasks":
			plan.Tasks = parseTasks(value.([]interface{}))
		}
	}

	return plan
}

// ParseTasks extracts all tasks from the given tasks (JSON representation).
func parseTasks(tasksSlice []interface{}) (tasks []tasks.Task) {
	for _, value := range tasksSlice {
		task := parseTask(value.(map[string]interface{}))
		tasks = append(tasks, task)
	}

	return tasks
}

// ParseTask creates a single task from the given task map (JSON representation).
func parseTask(taskMap map[string]interface{}) (task tasks.Task) {
	description := taskMap["Description"]
	if description == nil {
		description = ""
	}
	switch taskMap["Type"] {
	case "RenameFileTask":
		task = tasks.RenameFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "ExecCommandTask":
		task = tasks.ExecCommandTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "ServiceTask":
		task = tasks.ServiceTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "UnpackArchiveTask":
		task = tasks.UnpackArchiveTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "ConfirmationTask":
		task = tasks.ConfirmationTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "DeleteTask":
		task = tasks.DeleteTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "FindInFileTask":
		task = tasks.FindInFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	case "ReplaceInFileTask":
		task = tasks.ReplaceInFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: taskMap["Preferences"].(map[string]interface{}),
			},
		}
	}
	return task
}
