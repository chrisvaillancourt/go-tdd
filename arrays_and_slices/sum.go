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

func SumAll(numsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numsToSum {
		// append() takes a slice and a new value and returns a new slice with all items
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbers ...[]int) (sums []int) {
	for _, nums := range numbers {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
