package animals

// Cat defines an cat that implements the Humanizable interface
type Cat struct{}

// Path returns the URL path for cats
func (c *Cat) Path() string {
	return "cat"
}

// AgeInHumanYears calculates how many years given cat age is in human age
func (c *Cat) AgeInHumanYears(age int) int {
	return age * 4
}
