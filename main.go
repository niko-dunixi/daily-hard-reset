package main

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

//go:generate go run generate-cron.go
func main() {
	allApplications := getAllApplications()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(allApplications))
	for _, applicationName := range allApplications {
		go func(name string) {
			defer waitGroup.Done()
			cycleApplication(name)
		}(applicationName)
	}
	waitGroup.Wait()
}

func getAllApplications() AllApplications {
	hardResetFile, err := homedir.Expand("~/.daily-hard-reset.yaml")
	if err != nil {
		panic(err)
	}
	rawBytes, err := ioutil.ReadFile(hardResetFile)
	// Reasonable default if the file isn't there
	if os.IsNotExist(err) {
		rawBytes = []byte(`
- "Microsoft Outlook"
- "Slack"
- "zoom.us"
`)
	} else if err != nil {
		panic(err)
	}
	allApplications := AllApplications{}
	if err := yaml.Unmarshal(rawBytes, &allApplications); err != nil {
		panic(err)
	}
	return allApplications
}

type AllApplications []string

func cycleApplication(name string) {
	for stillRunning(name) {
		killApplication(name)
	}
	startApplication(name)
}

func killApplication(name string) {
	command := exec.Command("pkill", "-9", "-f", name)
	_ = command.Run()
}

func stillRunning(name string) bool {
	command := exec.Command("pgrep", name)
	// If this returns an error, the application is no longer running
	isRunning := command.Run() == nil
	return isRunning
}

func startApplication(name string) {
	command := exec.Command("open", "-a", name)
	_ = command.Run()
}
