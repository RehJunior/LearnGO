package calculator

import (
	"errors"
	"math"
)

func Abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func SquareRoot(value float64) (float64, error) {

	if value < 0 {
		return 0, errors.New("value nust not be negative")
	}
	return math.Sqrt(value), nil
}

func IsSquareNumber(value int) bool {
	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	} else {
		return true
	}
}
func RunOperation(opertaion string, left, right int) int {
	var result int
	switch op := opertaion; op {
	case "add":
		result = left + right
	case "substract":
		result = left - right
	default:
		// Intentionallly left blank
	}
	return result
}
