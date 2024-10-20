package goTests

import "testing"

func TestAdd(t *testing.T) {
	// arrange
	l, r := 1, 2
	expect := 3

	//act
	got := Add(l, r)

	//assert
	if got != expect {
		t.Errorf("Failed to add %v and %v, got %v, expected %v", l, r, got, expect)
	}
}
