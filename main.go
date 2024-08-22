package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/ollamaclient/v2"
)

const (
	model         = "gemma2:2b"
	prompt        = "Write a silly saying, quote or joke like it could have been the output of the fortune command on Linux. Only output the actual fortune."
	versionString = "fortune9000 v1.1.1"
)

func main() {
	versionFlag := flag.Bool("version", false, "Prints the version information")
	flag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		os.Exit(0)
	}

	oc := ollamaclient.New()
	oc.ModelName = model
	oc.SetRandom()

	if err := oc.PullIfNeeded(true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pull model: %v\n", err)
		os.Exit(1)
	}

	found, err := oc.Has(model)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not check for model: %v\n", err)
		os.Exit(1)
	}

	if !found {
		fmt.Fprintf(os.Stderr, "Expected to have the '%s' model downloaded, but it's not present\n", model)
		os.Exit(1)
	}

	oc.SetRandom()

	generatedOutput := oc.MustOutput(prompt)

	retryCounter := 0
	for len(generatedOutput) < 5 || (len(generatedOutput) > 1 && (generatedOutput[1] == ' ' || generatedOutput[1] == '.' || generatedOutput[1] == '-')) {
		generatedOutput = oc.MustOutput(prompt)
		retryCounter++
		if retryCounter > 7 || generatedOutput == "" {
			fmt.Println("I've got nothing.")
			return
		}
	}

	fmt.Println(strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(ollamaclient.Massage(generatedOutput), "```"), "```")))
}
