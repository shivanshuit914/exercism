package bob

import "strings"

const testVersion = 3

func Hey(input string) string {
	input = strings.TrimSpace(input)

	if input == "" {
		return "Fine. Be that way!"
	}

	if strings.ToUpper(input) == input && strings.ToLower(input) != input {
		return "Whoa, chill out!"
	}

	if (input[len(input)-1]) == '?' {
		return "Sure."
	}

	return "Whatever."
}
