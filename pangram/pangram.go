package pangram

import "strings"

const testVersion = 1

func IsPangram(input string) bool {
	isPan := true
	used := make(map[rune]int)
	input = strings.ToLower(input)
	for _, l := range input {
		used[l] = used[l] + 1
	}

	for i := 'a'; i <= 'z'; i++ {
		if used[i] == 0 {
			isPan = false
		}
	}
	return isPan
}
