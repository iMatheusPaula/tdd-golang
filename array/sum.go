package array

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumTraditionalFor(numbers []int) int {
	result := 0
	for i := 0; i < 5; i++ {
		result += numbers[i]
	}
	return result
}
