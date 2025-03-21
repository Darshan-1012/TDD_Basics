package arrndslices

func sum(numbers []int) int {
	sum := 0
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, sum(numbers))
	}
	return sums
}

func SumAllTrails(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, sum(tails))
		}
	}
	return sums
}
