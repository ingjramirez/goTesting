package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(arrays ...[]int) []int {
	var results []int
	for _, array := range arrays {
		results = append(results, Sum(array))
	}

	return results
}
