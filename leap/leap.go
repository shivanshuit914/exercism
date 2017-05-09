package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	var leapYear = false
	// Write some code here to pass the test suite.
	if year%4 == 0 {
		leapYear = true
	}

	if year%100 == 0 && year%400 != 0 {
		leapYear = false
	}

	return leapYear
}
