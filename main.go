// package main is the main package for the fortunecraft utility
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/xyproto/fullname"
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
	"github.com/xyproto/wordwrap"
	"golang.org/x/term"
)

const versionString = "FortuneCraft 1.8.1"

var prompt = "Write a clever saying, quote or joke that could have come from the fortune-mod application on Linux. Only output the fortune, in plain text."

// getTerminalWidth tries to find the current width of the terminal, with a fallback on 79
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 79 // fallback
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
	if maybeInappropriate && ((c(s, "request") && c(s, "fulfill")) || c(s, "appropriate") || c(s, "generating something different") || c(s, "conversation fun and safe") || c(s, "cannot provide content") || c(s, "something different") || c(s, "content") || c(s, "AI assist") || c(s, "for the purpose") || c(s, "isclaimer") || c(s, "offensive") || c(s, "ethical") || c(s, "'fortune") || c(s, "responsibly") || c(s, "suggestive")) {
		return true
	}
	return c(s, "apt-") || c(s, "apt ") || hp(s, ".") || hs(s, " a") || c(s, "et me know ") || c(s, "not be used to") || c(s, "simulated response") || c(s, "your instructions") || c(s, "interpreted as a statement")
}

// formatNicely tries to wrap the given string at the right places so that it looks good.
// This has been specially crafted for short fortunes that are often 1 or two sentences.
func formatNicely(s string) string {
	var (
		lineLen       = int(float64(getTerminalWidth()) * 0.6)
		lines, err    = wordwrap.WordWrap(s, lineLen)
		lastIndex     = len(lines) - 1
		filteredLines []string
	)
	for i, line := range lines {
		if strings.Contains(line, ". ") && !strings.HasSuffix(line, "..") {
			parts := strings.SplitN(line, ". ", 2)
			if len(parts[1]) < int(float64(lineLen)*0.35) {
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
		}
	}
	lastIndex = len(lines) - 1
	for len(lines) > 1 && len(lines[lastIndex]) < int(float64(lineLen)*0.1) {
		lastIndex = len(lines) - 1
		if lastIndex == 0 {
			break
		}
		lines[lastIndex-1] += " " + lines[lastIndex]
		lines = lines[:lastIndex]
		lastIndex = len(lines) - 1
	}
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
		fmt.Fprintln(os.Stderr, "  fortunecraft -I -k AI  - Generate ironic fortunes about AI")
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
