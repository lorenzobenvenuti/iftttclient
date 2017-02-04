package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lorenzobenvenuti/ifttt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("iftttclient", "A command-line IFTTT client application")

	triggerCmd    = app.Command("trigger", "Triggers an event")
	triggerApiKey = triggerCmd.Flag("api-key", "IFTTT API key").Short('k').String()
	triggerEvent  = triggerCmd.Arg("event", "Event to trigger").Required().String()
	triggerValues = triggerCmd.Arg("values", "Values to pass").Strings()

	storeCmd    = app.Command("store", "Stores the IFTTT API key for later use")
	storeApiKey = storeCmd.Arg("api-key", "IFTTT API key").Required().String()
)

func exit(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func getApiKeyFromStore() (string, error) {
	return NewStore("").Retrieve()
}

func getApiKey() (string, error) {
	key := *triggerApiKey
	if key == "" {
		key = os.Getenv("IFTTT_API_KEY")
		if key == "" {
			key, _ = getApiKeyFromStore()
		}
	}
	if key == "" {
		return "", errors.New("Please specify an API key")
	}
	return key, nil
}

func trigger() {
	key, err := getApiKey()
	if err != nil {
		exit(err)
	}
	iftttClient := ifttt.NewIftttClient(key)
	iftttClient.Trigger(*triggerEvent, *triggerValues)
}

func store() {
	NewStore("").Store(*storeApiKey)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			exit(err)
		}
	}()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case triggerCmd.FullCommand():
		trigger()
	case storeCmd.FullCommand():
		store()
	}

}
