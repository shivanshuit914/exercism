package hamming

import "errors"

const testVersion = 6

func Distance(dnaA, dnaB string) (int, error) {
	if len(dnaA) != len(dnaB) {
		return -1, errors.New("String length is not same.")
	}
	distance := 0
	for i := 0; i < len(dnaA); i++ {
		if dnaA[i] != dnaB[i] {
			distance++
		}
	}
	return distance, nil
}
