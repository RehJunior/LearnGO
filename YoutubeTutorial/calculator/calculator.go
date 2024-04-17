package calculator

//* geht mit datentyp oder ohne
//! := funktionoert nicht auf global ebene
var Hi = 3.141593654

const Pi float64 = 3.141593654

func Add(left, right int) int {
	//var sum int = left + right

	//var sum = left + right

	//var sum int
	//sum = left + right

	sum := left + right

	return sum

}
func Divide(left, right int) (quotient, remainder int) {
	quotient = left / right
	remainder = left % right
	return
}
