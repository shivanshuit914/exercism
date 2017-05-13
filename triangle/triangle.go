package triangle

import "math"

const testVersion = 3
const (
	NaT = iota
	Equ
	Iso
	Sca
)

type Kind int

func KindFromSides(a, b, c float64) Kind {
	if checkTriange(a, b, c) == false {
		return NaT
	}
	if a == b && a == c && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func checkTriange(a, b, c float64) bool {
	if a == 0 || b == 0 || c == 0 || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return (a+b >= c) && (a+c >= b) && (b+c >= a)
}
