// fullname/fullname.go
package fullname

import (
	"bufio"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/xyproto/env/v2"
)

// capitalizeFirst capitalizes the first letter of a given string
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

// Get attempts to retrieve the full name of the current user by checking multiple sources.
func Get() string {
	// Attempt to get the full name from the system user database
	if u, err := user.Current(); err == nil {
		if u.Name != "" && u.Name != u.Username {
			return u.Name
		}
	}

	// Check common environment variables for the full name
	// You can extend this list with more environment variables if needed
	fullName := env.StrAlt("FULLNAME", "USER_FULL_NAME")
	if fullName != "" {
		return fullName
	}

	// Parse the ~/.gitconfig file for the user.name field
	gitconfigPath := env.ExpandUser("~/.gitconfig")
	if contents, err := os.ReadFile(gitconfigPath); err == nil {
		scanner := bufio.NewScanner(strings.NewReader(string(contents)))
		inUserSection := false
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// Check for section headers
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				section := strings.ToLower(strings.Trim(line, "[]"))
				inUserSection = (section == "user")
				continue
			}

			// If within the [user] section, look for the name field
			if inUserSection && strings.HasPrefix(strings.ToLower(line), "name") {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					name := strings.TrimSpace(parts[1])
					if name != "" {
						return name
					}
				}
			}
		}
	}

	// Attempt to read from AccountsService
	// AccountsService stores user information in /var/lib/AccountsService/users/<username>
	accountsServicePath := filepath.Join("/var/lib/AccountsService/users", env.CurrentUser())
	if contents, err := os.ReadFile(accountsServicePath); err == nil {
		scanner := bufio.NewScanner(strings.NewReader(string(contents)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "FullName=") {
				fullName := strings.TrimPrefix(line, "FullName=")
				fullName = strings.TrimSpace(fullName)
				if fullName != "" {
					return fullName
				}
			}
		}
	}

	// Fallback: Use the username with the first letter capitalized
	userName := env.CurrentUser()
	if userName != "" {
		return capitalizeFirst(userName)
	}

	// If all methods fail, return an empty string
	return ""
}
