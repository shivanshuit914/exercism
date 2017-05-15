package diffsquares

const testVersion = 1

func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}

func SquareOfSums(number int) int {
	sum := 0
	for ; number > 0; number-- {
		sum += number
	}
	return square(sum)
}

func SumOfSquares(number int) int {
	sum := 0
	for ; number > 0; number-- {
		sum += square(number)
	}
	return sum
}

func square(number int) int {
	return number * number
}
