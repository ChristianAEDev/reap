//go:generate goversioninfo -icon=icon.ico -manifest=goversioninfo.exe.manifest
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ChristianAEDev/reap/actions"
	"github.com/ChristianAEDev/reap/config"
	"github.com/desertbit/grumble"
	log "github.com/sirupsen/logrus"
)

var app = grumble.New(&grumble.Config{
	Name:        "reap",
	Description: "Automate repetitiv tasks",
})

func main() {
	log.Info("Starting")
	grumble.Main(app)
	log.Info("Shutting down")
}

func init() {
	// Setup the logger to only log into a file and not stdout/stderr
	logFile, err := os.OpenFile("reapy.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(logFile)

	readConfig()
	app.AddCommand(&grumble.Command{
		Name: "list",
		Help: "List all available plans",
		Run:  actions.OnList,
	})

	execute := &grumble.Command{
		Name:      "execute",
		Help:      "Execute a choosen plan",
		Aliases:   []string{"exe"},
		AllowArgs: true,
		Flags: func(f *grumble.Flags) {
			f.Int("p", "plan", -1, "Plan to execute identified by it's index. (Use \"list\" to determine the index of the desired plan.")
		},
		Run: actions.OnExecute,
	}

	app.AddCommand(execute)
}

func readConfig() {
	// Read the plans from the plans.json file
	plansConfig, err := ioutil.ReadFile("plans.json")
	if err != nil {
		log.Panic(err)
	}
	config.AppConfig = config.ExtractConfig(plansConfig)
	fmt.Println("Plan(s) initialized:", len(config.AppConfig.Plans))
}
