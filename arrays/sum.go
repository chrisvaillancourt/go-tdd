package sum

func Sum(numbers []int) int {
	sum := 0
	// range lets us loop over the array
	// range returns the item's index & value
	for _, number := range numbers {
		sum += number
	}
	return sum
}
