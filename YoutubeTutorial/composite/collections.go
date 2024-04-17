package composite

import "fmt"

func DemoCollections() {
	primesArray := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primesArray)

	//primesSlice := []int{2, 3, 5, 7, 11}
	primesSlice := make([]int, 5, 7)
	primesSlice = append(primesSlice, 2)
	primesSlice = append(primesSlice, 3)
	primesSlice = append(primesSlice, 5)
	fmt.Println(primesSlice)
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))

	for _, value := range primesSlice {
		fmt.Println(value)
	}

	points := map[string]HasDistance{
		"A": NewPoint(3, 7),
		"B": NewPoint(2, 3),
	}
	fmt.Println(points)
	somePoint, ok := points["D"]
	fmt.Println(somePoint, ok)

	points["E"] = NewPoint(1, 1)
	delete(points, "A")
	fmt.Println(points)

}
