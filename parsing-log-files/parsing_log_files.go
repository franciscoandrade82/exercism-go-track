package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i).*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// Define the regex pattern to match "User " followed by a username
	re := regexp.MustCompile(`User\s+(\S+)`)

	// Create a slice to hold the modified lines
	var result []string

	// Iterate through each line
	for _, line := range lines {
		// Find the matches
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Extract the username from the matches
			username := matches[1]
			// Prefix the line with [USR] and the username
			newLine := "[USR] " + username + " " + line
			result = append(result, newLine)
		} else {
			// If no match is found, keep the line unchanged
			result = append(result, line)
		}
	}

	return result
}
