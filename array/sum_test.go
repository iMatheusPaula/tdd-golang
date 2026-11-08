package array

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	//O Go não te deixa usar operadores de igualdade com slices.
	//É possível escrever uma função que percorre cada slice recebido e esperado e verificar seus valores,
	//mas por praticidade podemos usar o reflect.DeepEqual que é útil para verificar se duas variáveis são iguais.
	verifySums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum multiple slices", func(t *testing.T) {
		result := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		verifySums(t, result, expected)
	})

	t.Run("sum empty slice", func(t *testing.T) {
		result := SumAll([]int{}, []int{0, 9})
		expected := []int{0, 9}

		verifySums(t, result, expected)
	})

	t.Run("sum multiple with append", func(t *testing.T) {
		result := SumAllWithAppend([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		verifySums(t, result, expected)
	})

	t.Run("sum slice, but not of the first number", func(t *testing.T) {
		result := SumFinalSlice([]int{10, 2, 10}, []int{9, 9, 8})
		expected := []int{12, 17}

		verifySums(t, result, expected)
	})

	t.Run("sum slice, but not of the last number", func(t *testing.T) {
		result := SumStartSlice([]int{7, 20, 1, 98}, []int{18, 900, 12, 77})
		expected := []int{28, 930}

		verifySums(t, result, expected)
	})

	t.Run("sum slice, empty slice", func(t *testing.T) {
		result := SumAll([]int{}, []int{0, 1})
		expected := []int{0, 1}

		verifySums(t, result, expected)
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
