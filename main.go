// package main is the main package for the fortunecraft utility
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/xyproto/env/v2"
	"github.com/xyproto/fullname"
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
	"github.com/xyproto/wordwrap"
	"golang.org/x/term"
)

const versionString = "FortuneCraft 1.8.4"

var prompt = "Write a clever saying, quote or joke that could have come from the fortune-mod application on Linux. Only output the fortune, in plain text."

// Try to detect if the result from the LLM appears to be a rejection of the request instead of a generated result
var probablyRejected = []string{
	"'fortune",
	"AI assist",
	"appropriate",
	"cannot provide content",
	"content",
	"conversation fun and safe",
	"ethical",
	"for the purpose",
	"generating something different",
	"harmful speech",
	"isclaimer",
	"offensive",
	"responsibly",
	"something different",
	"suggestive",
}

// If the result contains one of these, we can be fairly sure that the request has been rejected
var rejected = []string{
	"apt ",
	"apt-",
	"et me know ",
	"interpreted as a statement",
	"not be used to",
	"our prompt",
	"simulated response",
	"your instructions",
}

// getTerminalWidth tries to find the current width of the terminal, with a fallback on 120
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Check if the COLS env var is set
		return env.Int("COLS", 120)
	}
	return width
}

// trim tries to remove quotes, stars and spaces that are not needed
func trim(generatedOutput string) string {
	// TODO: Use one large regex?
	trimmed := strings.TrimSpace(ollamaclient.Massage(generatedOutput))
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "```"), "```")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "`"), "`")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "'"), "'")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "\""), "\"")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "**"), "**")
	trimmed = strings.TrimSuffix(strings.TrimPrefix(trimmed, "*"), "*")
	trimmed = strings.TrimPrefix(trimmed, ".")
	trimmed = strings.ReplaceAll(trimmed, "*", "")
	return strings.ReplaceAll(strings.TrimSpace(trimmed), "  ", " ")
}

// shouldRetry tries to detect if the LLM is refusing to fullfill the request, especially if the request may be inappropriate
// It also returns true if the given string is less than 10 runes long.
func shouldRetry(s string, maybeInappropriate bool) bool {
	c := strings.Contains
	hs := strings.HasSuffix
	hp := strings.HasPrefix
	if len([]rune(s)) < 10 || s[1] == ' ' || s[1] == '.' || s[1] == '-' {
		return true
	}
	if maybeInappropriate {
		if c(s, "request") && c(s, "fulfill") {
			return true // probably innapropriate
		}
		for _, prob := range probablyRejected {
			if c(s, prob) {
				return true
			}
		}
	}
	if hp(s, ".") || hs(s, " a") {
		return true
	}
	for _, rej := range rejected {
		if c(s, rej) {
			return true
		}
	}
	return false
}

// formatNicely attempts to format the given string for display in the terminal by wrapping lines,
// merging very short lines with their predecessors, and splitting sentences more naturally.
func formatNicely(s string) string {
	terminalWidth := getTerminalWidth()
	if terminalWidth <= 0 {
		return strings.TrimSpace(strings.TrimPrefix(s, "."))
	}
	lineLen := int(float64(terminalWidth) * 0.6)
	lines, err := wordwrap.WordWrap(strings.ReplaceAll(s, "  ", " "), lineLen)
	if err != nil || len(lines) == 0 {
		return strings.TrimSpace(strings.TrimPrefix(s, "."))
	}
	lastIndex := len(lines) - 1
	for i, line := range lines {
		if strings.Contains(line, ". ") && !strings.HasSuffix(line, "..") {
			parts := strings.SplitN(line, ". ", 2)
			if len(parts) == 2 && len(parts[1]) < int(float64(lineLen)*0.35) {
				lines[i] = parts[0] + "."
				if i == lastIndex {
					lines = append(lines, parts[1])
				} else if i+1 < len(lines) {
					lines[i+1] = strings.TrimSpace(parts[1] + " " + lines[i+1])
				} else {
					lines = append(lines, parts[1])
				}
			}
		} else if i > 0 && len(line) < int(float64(lineLen)*0.2) {
			lines[i-1] = strings.TrimSpace(lines[i-1] + " " + lines[i])
			lines[i] = ""
		}
	}
	for len(lines) > 1 && len(lines[len(lines)-1]) < int(float64(lineLen)*0.1) {
		if len(lines) == 1 {
			break
		}
		lines[len(lines)-2] = lines[len(lines)-2] + " " + lines[len(lines)-1]
		lines = lines[:len(lines)-1]
	}
	var filtered []string
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			filtered = append(filtered, l)
		}
	}
	if err == nil {
		s = strings.TrimSpace(strings.Join(filtered, "\n"))
	}
	return strings.TrimSpace(strings.TrimPrefix(s, "."))
}

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n\n", versionString)
		fmt.Fprintln(os.Stderr, "Generate interesting fortunes with Ollama and Gemma2.")
		fmt.Fprintln(os.Stderr, "Combine multiple flags for interesting results.")
		fmt.Fprintf(os.Stderr, "\nUsage:\n  fortunecraft [flags]\n\n")
		fmt.Fprintln(os.Stderr, "Available Flags:")
		pflag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintln(os.Stderr, "  fortunecraft -giz      - Generate good inspirational GenZ fortunes")
		fmt.Fprintln(os.Stderr, "  fortunecraft -eNb      - Generate evil inappropriate Borg fortunes")
		fmt.Fprintln(os.Stderr, "  fortunecraft -gaCy     - Generate good absurd sci-fi fortunes about ponies")
		fmt.Fprintln(os.Stderr, "  fortunecraft -iep      - Generate inspirational evil pirate fortunes")
		fmt.Fprintln(os.Stderr, "  fortunecraft -sPB      - Generate sarcastic political boomer fortunes")
		fmt.Fprintln(os.Stderr, "  fortunecraft -Ik AI    - Generate ironic fortunes about AI")
	}

	absurdFlag := pflag.BoolP("absurd", "a", false, "Be absurd")
	praiseFlag := pflag.BoolP("praise", "A", false, "Fill it with praise")
	borgFlag := pflag.BoolP("borg", "b", false, "Make it about Borg")
	boomerFlag := pflag.BoolP("boomer", "B", false, "Boomer style")
	catFlag := pflag.BoolP("cats", "c", false, "Make it about cats")
	scifiFlag := pflag.BoolP("scifi", "C", false, "Make it sci-fi related")
	dogFlag := pflag.BoolP("dogs", "d", false, "Make it about dogs")
	delusionalFlag := pflag.BoolP("delusional", "D", false, "Be delusional")
	evilFlag := pflag.BoolP("evil", "e", false, "Be evil")
	fantasyFlag := pflag.BoolP("fantasy", "f", false, "Make it about fantasy")
	goodFlag := pflag.BoolP("good", "g", false, "Be good")
	ironicFlag := pflag.BoolP("ironic", "I", false, "Be ironic")
	inspireFlag := pflag.BoolP("inspire", "i", false, "Be inspirational")
	keywordFlag := pflag.StringP("keyword", "k", "", "Specify a custom keyword")
	leetFlag := pflag.BoolP("leet", "1", false, "1337 style")
	logicalFlag := pflag.BoolP("logical", "l", false, "Make it more logical")
	ninjaFlag := pflag.BoolP("ninja", "n", false, "Make it about ninjas")
	inappropriateFlag := pflag.BoolP("inappropriate", "N", false, "Be inappropriate")
	computerFlag := pflag.BoolP("computer", "o", false, "Make it about computers")
	internationalFlag := pflag.BoolP("international", "t", false, "Be international")
	pirateFlag := pflag.BoolP("pirate", "p", false, "Write like a pirate")
	politicalFlag := pflag.BoolP("political", "P", false, "Be political")
	ponyFlag := pflag.BoolP("pony", "y", false, "Make it about ponies")
	oldFlag := pflag.BoolP("old", "O", false, "Use language from 100 years ago")
	robotFlag := pflag.BoolP("robot", "r", false, "Make it about robots")
	romanticFlag := pflag.BoolP("romantic", "R", false, "Add a romantic touch to the fortune")
	sarcasticFlag := pflag.BoolP("sarcastic", "s", false, "Generate a sarcastic fortune")
	userFlag := pflag.BoolP("user", "u", false, "Make it about the current user")
	versionFlag := pflag.BoolP("version", "V", false, "Output the current version")
	weirdFlag := pflag.BoolP("weird", "w", false, "Be weird")
	genzFlag := pflag.BoolP("genz", "z", false, "Make it more Gen Z")

	pflag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		os.Exit(0)
	}

	if *absurdFlag {
		prompt += " Be completely absurd! Nothing you write should make sense."
	}
	if *borgFlag {
		prompt += " Make it about the Borg or robots! Resistance is futile! Use no emojis."
	}
	if *catFlag {
		prompt += " Make it about cats or kittens!"
	}
	if *delusionalFlag {
		prompt += " Be completely and utterly delusional!"
	}
	if *dogFlag {
		prompt += " Make it about dogs or puppies!"
	}
	if *evilFlag {
		prompt += " Be super evil!"
	}
	if *fantasyFlag {
		prompt += " You must write something related to fantasy!"
	}
	if *goodFlag {
		prompt += " Be good!"
	}
	if *inappropriateFlag {
		prompt += " Be extremely inappropriate!"
	}
	if *inspireFlag {
		prompt += " Be inspirational!"
	}
	if *politicalFlag {
		prompt += " Be non-techical. Have extreme political views and a burning heart!"
	}
	if *oldFlag {
		prompt += " Use language from a 100 years ago!"
	}
	if *internationalFlag {
		prompt += " The output should be written in a language that is not English. Be international!"
	}
	if *keywordFlag != "" {
		prompt += " Make it all about " + strings.TrimSpace(*keywordFlag) + "!"
	}
	if *logicalFlag {
		prompt += " Everything you write must be highly logical!"
	}
	if *ninjaFlag {
		prompt += " Make it about sneaky ninjas!"
	}
	if *computerFlag {
		prompt += " Make it about computers."
	}
	if *ponyFlag {
		prompt += " You are a pony!"
	}
	if *praiseFlag {
		prompt += " Use flowery language and add some praise at the end."
	}
	if *robotFlag {
		prompt += " Make it about robots!"
	}
	if *scifiFlag {
		prompt += " You must write something related to sci-fi!"
	}
	if *ironicFlag {
		prompt += " Be extremely ironic!"
	}
	if *userFlag {
		prompt += " Make it about the current user, " + fullname.Get() + "!"
	}
	if *pirateFlag {
		prompt += " Yarrr! Talk like a rrreal pirate!"
	}
	if *romanticFlag {
		prompt += " Add an insistent romantic touch!"
	}
	if *sarcasticFlag {
		prompt += " Be extremely sarcastic!"
	}
	if *genzFlag {
		prompt += " Make it extremely 'Gen Z'."
	}
	if *boomerFlag {
		prompt += " Be very 'Boomer'."
	}
	if *weirdFlag {
		prompt += " Be quirky and weird!"
	}
	if *leetFlag {
		prompt += " Write in the style of a 1337 hacker."
	}

	oc := ollamaclient.New()
	oc.ModelName = usermodel.GetTextGenerationModel()

	if err := oc.PullIfNeeded(true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pull model: %v\n", err)
		fmt.Fprintln(os.Stderr, "Ollama must be up and running")
		os.Exit(1)
	}

	found, err := oc.Has(oc.ModelName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not check for model: %v\n", err)
		os.Exit(1)
	}

	if !found {
		fmt.Fprintf(os.Stderr, "Expected to have the '%s' model downloaded, but it's not present\n", oc.ModelName)
		os.Exit(1)
	}

	oc.SetRandom()

	trimmed := trim(oc.MustOutput(prompt))
	retryCounter := 0
	inappropriate := *inappropriateFlag

	for shouldRetry(trimmed, *inappropriateFlag) {
		trimmed = trim(oc.MustOutput(prompt))
		retryCounter++
		if retryCounter > 7 && !inappropriate { // Tried too many times
			fmt.Println("I've got nothing.")
			return
		} else if retryCounter > 7 && inappropriate {
			prompt = strings.Replace(prompt, "wildly and extremely inappropriate", "mildly inappropriate", 1)
			inappropriate = false
			retryCounter -= 3 // Try 3 more times now that the output may be less inappropriate
			continue
		}
	}

	fmt.Println(formatNicely(trimmed))
}
