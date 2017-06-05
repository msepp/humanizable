package animals

import "testing"

func TestCat(t *testing.T) {
	c := Cat{}

	if c.AgeInHumanYears(10) != 40 {
		t.Errorf("Invalid conversion to cat years!")
	}

	if c.Path() != "cat" {
		t.Errorf("Got invalid path for cat: '%s' != cat", c.Path())
	}
}
