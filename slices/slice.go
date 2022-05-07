package slices

//Sum returns the sum of the elements in an array of size 4.
func Sum(array []int) int {
	var sum int

	for _, value := range array {
		sum += value
	}
	return sum
}

//SumAll get a variable number of int slices and returns a slice with the sum
// of the elements of those slices.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//SumAllTails same behavior of SumAll but sums every element except the first one (head).
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}

	}
	return sums
}
