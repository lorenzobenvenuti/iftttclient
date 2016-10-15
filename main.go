package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lorenzobenvenuti/ifttt"
)

type valueFlags []string

func (i *valueFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *valueFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func exit(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

var apiKey string
var event string
var iftttValues valueFlags

func checkApiKey() {
	if apiKey == "" {
		apiKey = os.Getenv("IFTTT_API_KEY")
		if apiKey == "" {
			exit("Please specify an API key or set the IFTTT_API_KEY variable")
		}
	}
}

func checkEvent() {
	if event == "" {
		exit("Please specify an event to trigger")
	}
}

func checkArguments() {
	checkApiKey()
	checkEvent()
}

func parseCommandLineArguments() {
	flag.StringVar(&apiKey, "key", "", "IFTTT API key")
	flag.StringVar(&event, "event", "", "Event name")
	flag.Var(&iftttValues, "value", "A value")
	flag.Parse()
}

func trigger() {
	iftttClient := ifttt.NewIftttClient(apiKey)
	iftttClient.Trigger(event, iftttValues)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			exit(err)
		}
	}()
	parseCommandLineArguments()
	checkArguments()
	trigger()
}
