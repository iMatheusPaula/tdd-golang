package interaction

import "testing"

func TestInteraction(t *testing.T) {
	interaction := Interaction("a")
	expect := "aaaaaa"

	if interaction != expect {
		t.Errorf("expect %s, got %s", expect, interaction)
	}
}

func BenchmarkInteraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interaction("a")
	}
}
