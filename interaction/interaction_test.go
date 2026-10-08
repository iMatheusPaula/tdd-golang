package interaction

import (
	"fmt"
	"strings"
	"testing"
)

func TestInteraction(t *testing.T) {
	interaction := Interaction("a", 6)
	expect := "aaaaaa"

	if interaction != expect {
		t.Errorf("expect %s, got %s", expect, interaction)
	}
}

func BenchmarkInteraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interaction("a", 10)
	}
}

func BenchmarkNativeRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 10)
	}
}

func ExampleInteraction() {
	result := Interaction("a", 6)
	fmt.Println(result)
	//Output: aaaaaa
}
