// Package actions contains the implementation for the functions the CLI tool provides.
package actions

import (
	"fmt"
	"strconv"

	"github.com/ChristianAEDev/reap/config"
	"github.com/desertbit/grumble"
)

// OnExecutes iterates over all task of a given plan and executes them.
func OnExecute(c *grumble.Context) error {
	planID := c.Flags.Int("plan")
	if len(config.AppConfig.Plans) < planID {
		fmt.Println("No plan with ID", planID, "exists. Use the command \"list\" to get a list of available plans.")
		return nil
	}
	plan := config.AppConfig.Plans[planID-1]

	for _, task := range plan.Tasks {
		result := task.Execute()

		if result.IsSuccessful {

			fmt.Println(result.Message)
		} else {
			fmt.Println("Error:", result.Message)
		}
	}

	return nil
}

// OnList prints all available plans including their tasks to the ommand line.
func OnList(c *grumble.Context) error {
	len := len(config.AppConfig.Plans)
	if len == 0 {
		fmt.Println("No plans available. Try \"read\".")
		return nil
	}

	for i, plan := range config.AppConfig.Plans {
		fmt.Println(i+1, plan.Name)
		for i, task := range plan.Tasks {
			if task.GetDescription() != "" {
				fmt.Println("  ", strconv.Itoa(i+1)+".", task.GetDescription())
			}
		}
	}
	return nil
}
