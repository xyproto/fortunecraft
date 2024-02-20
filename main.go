package main

import (
	"fmt"
	"os"

	"github.com/xyproto/ollamaclient"
)

const (
	model  = "tinydolphin"
	prompt = "Write a silly saying, quote or joke like it could have been the output of the fortune command on Linux."
)

func main() {
	oc := ollamaclient.NewWithModel(model)

	if err := oc.PullIfNeeded(true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pull model: %v\n", err)
		os.Exit(1)
	}

	found, err := oc.Has2(model)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not check for model: %v\n", err)
		os.Exit(1)
	}

	if !found {
		fmt.Fprintf(os.Stderr, "Expected to have the '%s' model downloaded, but it's not present\n", model)
		os.Exit(1)
	}

	oc.SetRandomOutput()

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

	fmt.Println(ollamaclient.Massage(generatedOutput))
}
