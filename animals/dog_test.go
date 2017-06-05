package animals

import "testing"

func TestDog(t *testing.T) {
	d := Dog{}

	if d.AgeInHumanYears(10) != 50 {
		t.Errorf("Invalid conversion to dog years!")
	}

	if d.Path() != "dog" {
		t.Errorf("Got invalid path for dog: '%s' != dog", d.Path())
	}
}
