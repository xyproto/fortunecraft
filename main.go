package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/xyproto/fullname"
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/wordwrap"
	"golang.org/x/term"
)

const (
	versionString    = "fortuna 1.4.0"
	model            = "gemma2:2b"
	defaultTermWidth = 79
)

var prompt = "Write a silly saying, quote or joke, like it could have been the output of the fortune-mod program on Linux. Only output the actual fortune, without emojis, using plain text."

func trim(generatedOutput string) string {
	trimmed := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(ollamaclient.Massage(generatedOutput), "```"), "```"))
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "'"), "'")
	return strings.ReplaceAll(strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(trimmed, "`"), "`")), "  ", " ")
}

func shouldRetry(s string) bool {
	return len(s) < 10 || (len(s) > 1 && (s[1] == ' ' || s[1] == '.' || s[1] == '-')) || strings.Contains(s, "apt-") || strings.Contains(s, " apt ") || strings.HasPrefix(s, ".")
}

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return defaultTermWidth
	}
	return width
}

func formatNicely(s string) string {
	lineLen := int(float64(getTerminalWidth()) * 0.7)
	lines, err := wordwrap.WordWrap(s, lineLen)
	lastIndex := len(lines) - 1
	var skipList []int
	for i, line := range lines {
		if strings.Contains(line, ". ") {
			parts := strings.SplitN(line, ". ", 2)
			if len(parts[1]) < int(float64(lineLen)*0.3) {
				lines[i] = parts[0] + "."
				if i == lastIndex {
					lines = append(lines, parts[1])
				} else {
					lines[i+1] = strings.TrimSpace(parts[1] + " " + lines[i+1])
				}
			}
		} else if i > 0 && len(line) < int(float64(lineLen)*0.2) {
			lines[i-1] = strings.TrimSpace(lines[i-1] + " " + lines[i])
			lines[i] = ""
			skipList = append(skipList, i)
		}
	}
	var filteredLines []string
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	if err == nil { // success
		s = strings.TrimSpace(strings.Join(filteredLines, "\n"))
	}
	return s
}

func main() {
	absurdFlag := pflag.BoolP("absurd", "a", false, "Be absurd")
	catFlag := pflag.BoolP("cats", "c", false, "Make it about cats")
	dogFlag := pflag.BoolP("dogs", "d", false, "Make it about dogs")
	evilFlag := pflag.BoolP("evil", "e", false, "Be evil")
	fantasyFlag := pflag.BoolP("fantasy", "y", false, "Make it about fantasy")
	goodFlag := pflag.BoolP("good", "g", false, "Be good")
	hotFlag := pflag.BoolP("hot", "h", false, "Be hot")
	internationalFlag := pflag.BoolP("international", "i", false, "Be international")
	logicalFlag := pflag.BoolP("logical", "l", false, "Make it more logical")
	ninjaFlag := pflag.BoolP("ninja", "n", false, "Make it about ninjas")
	politicalFlag := pflag.BoolP("political", "o", false, "Be political")
	ponyFlag := pflag.BoolP("pony", "p", false, "Make it about ponies")
	robotFlag := pflag.BoolP("robot", "r", false, "Make it about robots")
	scifiFlag := pflag.BoolP("scifi", "f", false, "Make it sci-fi related")
	snarkyFlag := pflag.BoolP("snarky", "s", false, "Be snarky")
	userFlag := pflag.BoolP("user", "u", false, "Make it about the current user")
	versionFlag := pflag.BoolP("version", "v", false, "Output the current version")
	weirdFlag := pflag.BoolP("weird", "w", false, "Be weird")

	pflag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		os.Exit(0)
	}

	if *snarkyFlag {
		prompt += " Be *extremely* snarky, sassy, pedantic and/or sarcastic!"
	}

	if *politicalFlag {
		prompt += " Be political!"
	}

	if *goodFlag {
		prompt += " Be good!"
	}

	if *evilFlag {
		prompt += " Be evil!"
	}

	if *scifiFlag {
		prompt += " You must write something related to sci-fi!"
	}

	if *internationalFlag {
		prompt += " The output should be written in a language that is not English. Be international!"
	}

	if *absurdFlag {
		prompt += " Be completely absurd and absurd! Nothing you write should make sense."
	}

	if *fantasyFlag {
		prompt += " You must write something related to fantasy!"
	}

	if *logicalFlag {
		prompt += " Everything you write must be highly logical!"
	}

	if *userFlag {
		prompt += " Praise the current user, " + fullname.Get() + "!"
	}

	if *hotFlag {
		prompt += " Be hot!"
	}

	if *weirdFlag {
		prompt += " Be quirky and weird!"
	}

	if *catFlag {
		prompt += " Make it about cats or kittens!"
	}

	if *dogFlag {
		prompt += " Make it about dogs or puppies!"
	}

	if *ponyFlag {
		prompt += " Make it about ponies!"
	}

	if *robotFlag {
		prompt += " Make it about robots!"
	}

	if *ninjaFlag {
		prompt += " Make it about ninjas!"
	}

	oc := ollamaclient.New()
	oc.ModelName = model

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

	trimmed := trim(oc.MustOutput(prompt))
	retryCounter := 0
	for shouldRetry(trimmed) {
		trimmed = trim(oc.MustOutput(prompt))
		retryCounter++
		if retryCounter > 7 || trimmed == "" {
			fmt.Println("I've got nothing.")
			return
		}
	}

	trimmed = formatNicely(trimmed)

	fmt.Println(trimmed)
}
