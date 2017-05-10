package acronym

import (
	"log"
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(inputString string) string {
	var returnstring string

	inputString = strings.Title(inputString)

	reg, err := regexp.Compile("[a-z0-9,:-]")
	if err != nil {
		log.Fatal(err)
	}
	inputString = strings.Replace(inputString, "-", " ", 1)
	inputString = reg.ReplaceAllString(inputString, "")
	splitted := strings.Split(inputString, " ")
	for _, value := range splitted {
		returnstring = returnstring + value[0:1]
	}

	return returnstring
}
