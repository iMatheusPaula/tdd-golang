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

func SumAll(numbers ...[]int) (sums []int) {
	quantityOfNumbers := len(numbers)
	sums = make([]int, quantityOfNumbers)

	for i, number := range numbers {
		sums[i] = Sum(number)
	}

	return
}

func SumAllWithAppend(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return
}

func SumFinalSlice(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		final := number[1:]
		sums = append(sums, Sum(final))
	}

	return
}

func SumStartSlice(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		start := number[:len(number)-1]
		sums = append(sums, Sum(start))
	}

	return
}
