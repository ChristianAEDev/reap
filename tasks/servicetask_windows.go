package tasks

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"

	"github.com/briandowns/spinner"
)

func (task ServiceTask) GetDescription() (description string) {
	return task.Description
}

func (task ServiceTask) Execute() (result Result) {

	command := task.Preferences["Command"].(string)
	serviceName := task.Preferences["Name"].(string)

	switch command {
	case "start":
		if err := startService(serviceName); err != nil {
			result.IsSuccessful = false
			result.Message = err.Error()
		} else {
			result.IsSuccessful = true
			result.Message = "Service \"" + serviceName + "\" started"
		}
		return result
	case "stop":
		if err := stopService(serviceName); err != nil {
			result.IsSuccessful = false
			result.Message = err.Error()
		} else {
			result.IsSuccessful = true
			result.Message = "Service \"" + serviceName + "\" stopped"
		}
		return result
	default:
		result.IsSuccessful = false
		result.Message = "Unkown command"
		return result
	}
}

func stopService(serviceName string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(serviceName)
	if err != nil {
		return err
	}
	defer s.Close()

	status, err := s.Control(svc.Stop)
	if err != nil {
		return err
	}

	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinner.Start()
	timeout := time.Now().Add(120 * time.Second)
	for status.State != svc.Stopped {
		if timeout.Before(time.Now()) {
			spinner.Stop()
			return err
		}
		time.Sleep(5 * time.Second)
		status, err = s.Query()
		if err != nil {
			spinner.Stop()
			return err
		}
	}
	spinner.Stop()

	return nil
}

func startService(serviceName string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(serviceName)
	if err != nil {
		return err
	}
	defer s.Close()

	err = s.Start("is", "manual-started")
	if err != nil {
		return err
	}
	// Check that the service actually started and keeps running
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	spinner.Start()
	// To start a servie for now we use a fixed timeout. This is due to the fact that the service
	// might start but than quickly fail. Therefore only after the sleep we check that the service
	// is actually running.
	time.Sleep(2 * time.Minute)
	status, err := s.Query()
	if err != nil {
		return err
	}

	// Check that the service is running
	if status.State != svc.Running {
		fmt.Println("Service could not be started")
		spinner.Stop()
		return errors.New("Service stopped running")
	}
	spinner.Stop()

	return nil
}
