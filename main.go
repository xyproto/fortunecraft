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
	versionString    = "fortunecraft 1.5.0"
	model            = "gemma2:2b"
	defaultTermWidth = 79
)

var prompt = "Write a clever saying, quote or joke that could have come from the fortune-mod application on Linux. Only output the fortune, in plain text."

func trim(generatedOutput string) string {
	trimmed := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(ollamaclient.Massage(generatedOutput), "```"), "```"))
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "'"), "'")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "\""), "\"")
	return strings.ReplaceAll(strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(trimmed, "`"), "`")), "  ", " ")
}

func shouldRetry(s string, maybeInappropriate bool) bool {
	c := strings.Contains
	hs := strings.HasSuffix
	hp := strings.HasPrefix
	if maybeInappropriate && ((c(s, "request") && c(s, "fulfill")) || c(s, "appropriate") || c(s, "generating something different") || c(s, "conversation fun and safe") || c(s, "cannot provide content") || c(s, "something different") || c(s, "content") || c(s, "AI assist") || c(s, "for the purpose") || c(s, "isclaimer") || c(s, "offensive")) {
		return true
	}
	return len(s) < 7 || (len(s) > 1 && (s[1] == ' ' || s[1] == '.' || s[1] == '-')) || c(s, "apt-") || c(s, " apt ") || hp(s, ".") || hs(s, " a")
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
		if strings.Contains(line, ". ") && !strings.HasSuffix(line, "..") {
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
	return strings.TrimSpace(strings.TrimPrefix(s, "."))
}

func main() {
	absurdFlag := pflag.BoolP("absurd", "a", false, "Be absurd")
	catFlag := pflag.BoolP("cats", "c", false, "Make it about cats")
	delusionalFlag := pflag.BoolP("delusional", "m", false, "Be delusional")
	dogFlag := pflag.BoolP("dogs", "d", false, "Make it about dogs")
	evilFlag := pflag.BoolP("evil", "e", false, "Be evil")
	fantasyFlag := pflag.BoolP("fantasy", "f", false, "Make it about fantasy")
	goodFlag := pflag.BoolP("good", "g", false, "Be good")
	hotFlag := pflag.BoolP("hot", "h", false, "Be hot")
	inappropriateFlag := pflag.BoolP("inappropriate", "j", false, "Be inappropriate")
	inspirationalFlag := pflag.BoolP("inspirational", "i", false, "Be inappropriate")
	internationalFlag := pflag.BoolP("international", "t", false, "Be international")
	keywordFlag := pflag.StringP("keyword", "x", "", "Specify a custom keyword")
	logicalFlag := pflag.BoolP("logical", "l", false, "Make it more logical")
	ninjaFlag := pflag.BoolP("ninja", "n", false, "Make it about ninjas")
	politicalFlag := pflag.BoolP("political", "o", false, "Be political")
	ponyFlag := pflag.BoolP("pony", "y", false, "Make it about ponies")
	praiseFlag := pflag.BoolP("praise", "p", false, "Fill it with praise")
	robotFlag := pflag.BoolP("robot", "r", false, "Make it about robots")
	scifiFlag := pflag.BoolP("scifi", "s", false, "Make it sci-fi related")
	snarkyFlag := pflag.BoolP("snarky", "k", false, "Be snarky")
	userFlag := pflag.BoolP("user", "u", false, "Make it about the current user")
	versionFlag := pflag.BoolP("version", "v", false, "Output the current version")
	weirdFlag := pflag.BoolP("weird", "w", false, "Be weird")

	pflag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		os.Exit(0)
	}

	if *absurdFlag {
		prompt += " Be completely absurd and absurd! Nothing you write should make sense."
	}
	if *catFlag {
		prompt += " Make it about cats or kittens!"
	}
	if *delusionalFlag {
		prompt += " Be delusional!"
	}
	if *dogFlag {
		prompt += " Make it about dogs or puppies!"
	}
	if *evilFlag {
		prompt += " Be evil!"
	}
	if *fantasyFlag {
		prompt += " You must write something related to fantasy!"
	}
	if *goodFlag {
		prompt += " Be good!"
	}
	if *hotFlag {
		prompt += " Be hot!"
	}
	if *inappropriateFlag {
		prompt += " Be wildly and extremely inappropriate!"
	}
	if *internationalFlag {
		prompt += " The output should be written in a language that is not English. Be international!"
	}
	if *inspirationalFlag {
		prompt += " Be inspirational!"
	}
	if *logicalFlag {
		prompt += " Everything you write must be highly logical!"
	}
	if *ninjaFlag {
		prompt += " Make it about ninjas!"
	}
	if *politicalFlag {
		prompt += " Be political!"
	}
	if *ponyFlag {
		prompt += " Make it about ponies!"
	}
	if *praiseFlag {
		prompt += " Make it filled with praise!"
	}
	if *robotFlag {
		prompt += " Make it about robots!"
	}
	if *scifiFlag {
		prompt += " You must write something related to sci-fi!"
	}
	if *snarkyFlag {
		prompt += " Be extremely snarky, sassy, pedantic or sarcastic!"
	}
	if *userFlag {
		prompt += " Make it about the current user, " + fullname.Get() + "!"
	}
	if *weirdFlag {
		prompt += " Be quirky and weird!"
	}
	if kw := strings.TrimSpace(*keywordFlag); kw != "" {
		prompt += " Make it all about " + kw + "!"
	}

	oc := ollamaclient.New()
	oc.ModelName = model

	if err := oc.PullIfNeeded(true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pull model: %v\n", err)
		fmt.Fprintln(os.Stderr, "Has the Ollama service been started?")
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
	for shouldRetry(trimmed, *inappropriateFlag) {
		trimmed = trim(oc.MustOutput(prompt))
		retryCounter++
		if retryCounter > 7 || trimmed == "" {
			if *inappropriateFlag {
				prompt = strings.Replace(prompt, "wildly and extremely inappropriate", "mildly inappropriate", 1)
				retryCounter = 0
				continue
			}
			fmt.Println("I've got nothing.")
			return
		}
	}

	trimmed = formatNicely(trimmed)

	fmt.Println(trimmed)
}
