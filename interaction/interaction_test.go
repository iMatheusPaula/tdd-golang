package interaction

import "testing"

func TestInteraction(t *testing.T) {
	interaction := Interaction("a")
	expect := "aaaaaa"

	if interaction != expect {
		t.Errorf("expect %s, got %s", expect, interaction)
	}
}
