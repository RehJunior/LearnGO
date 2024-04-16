package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}
func Sum(a, b, c, d, e int) int {
	sum := a + b + c + d + e
	return sum
}
