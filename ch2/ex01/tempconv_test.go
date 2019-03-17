package tempconv

import (
	"testing"
)

func TestExercise(t *testing.T) {
	c := Celsius(-273.15)
	k := CToK(c)

	if k != 0 {
		t.Fatalf("%v", k)
	}
}
