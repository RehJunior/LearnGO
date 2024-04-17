package main

import (
	"fmt"
	"time"

	"github.com/RehJunior/calculator"
	"github.com/RehJunior/composite"
)

func Print(value interface{}) {

	switch value.(type) {
	case float64:
		fmt.Println(value)
	default:
		fmt.Println("Can not be printed Manyuuk you crashed the Terminal Stupid")
	}
}

func main() {
	Print("Hallo Rene")
	time.Sleep(time.Second * 10)
	queue := make(chan int, 100)

	go func(q chan int) {
		time.Sleep(time.Second * 20)
		fmt.Println("Ok ich mach Spaß gönn dir dein Ergebis Habibi")
		q <- 23
	}(queue)

	fmt.Println("23 + 42 = ", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(17, 3))
	fmt.Println("Dachtest wohl geht weiter Manyuuk")

	time.Sleep(time.Second * 15)
	valueFromQueue := <-queue

	fmt.Println(calculator.Sum(1, valueFromQueue))

	fmt.Println(calculator.SumUntil(1000))
	//calculator.SumInfinite()
	fmt.Println(calculator.IsSquareNumber(25))

	fmt.Println(calculator.MultiplyFromAToB(1, 10))

	fmt.Println(composite.Add(23, 42))

	//pointPrt := &point
	point := composite.NewPoint(3, 7)
	//fmt.Println(point.(composite.Point).X, point.(composite.Point).Y)
	fmt.Println(point.X, point.Y)

	//composite.DemoCollections()

	Print(point.DistanceFromZero())
	fmt.Println(point)

	result, err := calculator.SquareRoot(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
