// Package config holds functions responsible to handle a config file. The main functionallity
// is to create the correct tasks.
package config

import (
	"encoding/json"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"

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
	// Values can contain global variables that are used througout the rest of the plan.
	Variables map[string]string
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
		case "Variables":
			plan.Variables = parseValues(value.(map[string]interface{}))
		case "Tasks":
			plan.Tasks = parseTasks(value.([]interface{}), plan.Variables)
		}
	}
	return plan
}

func parseValues(variablesMap map[string]interface{}) (variables map[string]string) {
	variables = map[string]string{}
	for key, value := range variablesMap {
		variables[key] = value.(string)
	}
	return variables
}

// ParseTasks extracts all tasks from the given tasks (JSON representation).
func parseTasks(tasksSlice []interface{}, variables map[string]string) (tasks []tasks.Task) {
	for _, value := range tasksSlice {
		task := parseTask(value.(map[string]interface{}), variables)
		tasks = append(tasks, task)
	}

	return tasks
}

// ParseTask creates a single task from the given task map (JSON representation).
func parseTask(taskMap map[string]interface{}, variables map[string]string) (task tasks.Task) {
	description := taskMap["Description"]
	if description == nil {
		description = ""
	}
	switch taskMap["Type"] {
	case "RenameFileTask":
		task = tasks.RenameFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "ExecCommandTask":
		task = tasks.ExecCommandTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "ServiceTask":
		task = tasks.ServiceTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "UnpackArchiveTask":
		task = tasks.UnpackArchiveTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "ConfirmationTask":
		task = tasks.ConfirmationTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "DeleteTask":
		task = tasks.DeleteTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "FindInFileTask":
		task = tasks.FindInFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	case "ReplaceInFileTask":
		task = tasks.ReplaceInFileTask{
			RunnableTask: tasks.RunnableTask{
				Description: description.(string),
				Preferences: expandPreferences(taskMap["Preferences"].(map[string]interface{}), variables),
			},
		}
	}
	return task
}

// expandPreferences checks each preference if it contains a variable "${...}" if it finds one
// it will check that this variable has been defined and fills it accordingly. If a string contains
// a variable that is not defined an error is returned.
func expandPreferences(preferences map[string]interface{}, variables map[string]string) (expandedPreferences map[string]interface{}) {
	expandedPreferences = make(map[string]interface{})
	for key, value := range preferences {
		// A slice has to be handles different thatn a map
		switch value.(type) {
		default:
			log.Error("Unknown preference type in config")
		case []interface{}:
			// For a JSON array
			for k, v := range value.([]interface{}) {
				value.([]interface{})[k] = expandVariable(v.(string), variables)
			}
			expandedPreferences[key] = value
		case interface{}:
			// For a simple field in JSON
			expandedPreferences[key] = expandVariable(value.(string), variables)
		}
	}
	return expandedPreferences
}

func expandVariable(preference string, variables map[string]string) (expandedPreference string) {
	expandedPreference = preference
	// See if there are variables in the preference
	r, _ := regexp.Compile(`\${(.*?)\}`)
	hits := r.FindAllStringIndex(preference, -1)
	// Iterate over all hits
	for _, hit := range hits {
		h := preference[hit[0]:hit[1]]
		// Strip the beginning "${" and end "}" of the variable to get it's name
		variableName := h[2 : len(h)-1]
		// Do we have a variables with that name
		variableValue := variables[variableName]
		expandedPreference = strings.Replace(preference, h, variableValue, -1)
	}
	return expandedPreference
}
