package raindrops

import "fmt"

const testVersion = 3

var rain = make([]string, 10)

func Convert(inputNumber int) string {
	sounds := ""
	rain[3] = "Pling"
	rain[5] = "Plang"
	rain[7] = "Plong"
	for i, sound := range rain {
		if 0 < i && inputNumber%i == 0 {
			sounds += sound
		}
	}
	if sounds == "" {
		sounds = fmt.Sprintf("%d", inputNumber)
	}
	return sounds
}
