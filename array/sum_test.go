package array

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("Sum(%d) = %d; expected %d", numbers, result, expected)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum([5]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumTraditionalFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumTraditionalFor([5]int{1, 2, 3, 4, 5})
	}
}
