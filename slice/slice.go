package slice

func Construction() []int {
	slice := make([]int, 10)
	for i, _ := range slice {
		slice[i] = 3
	}
	return slice
}

func Map() []int {
	slice := []int{1, 2, 3}
	for i, x := range slice {
		slice[i] = x * 2
	}
	return slice
}
