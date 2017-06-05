package main

// Humanizable interface tells how to humanize animals
type Humanizable interface {
	// Path should return an URL path for accessing the humanizable animal
	Path() string
	// AgeInHumanYears should produce a conversion of animal age into comparable human age
	AgeInHumanYears(int) int
}
