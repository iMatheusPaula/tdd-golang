package array

import "testing"

func TestSum(t *testing.T) {
	verifyCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("five numbers slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if result != expected {
			verifyCorrectMessage(t, result, expected)
		}
	})

	t.Run("any length slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 10, 1000}
		result := Sum(numbers)
		expected := 1025

		if result != expected {
			verifyCorrectMessage(t, result, expected)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumTraditionalFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumTraditionalFor([]int{1, 2, 3, 4, 5})
	}
}
