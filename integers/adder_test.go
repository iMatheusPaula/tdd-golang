package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got: %d, want: %d.", sum, expected)
	}
}

func ExampleAdder() {
	sum := Adder(2, 5)
	fmt.Println(sum)
	//Output: 7
}
