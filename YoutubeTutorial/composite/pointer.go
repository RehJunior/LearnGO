package composite

func Add(left, right int) int {
	//var leftPtr *int = &left
	leftPtr := &left
	rightPtr := &right

	sum := *leftPtr + *rightPtr

	return sum
}
