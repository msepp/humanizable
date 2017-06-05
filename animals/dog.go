package animals

// Dog defines a dog that implements the Humanizable interface
type Dog struct{}

// Path returns the URL path for dogs
func (d *Dog) Path() string {
	return "dog"
}

// AgeInHumanYears calculates how many years given dog age is in human age
func (d *Dog) AgeInHumanYears(age int) int {
	return age * 5
}
