// dog is all about dogs.
package dog

// Dog is a representation of a dog.
type Dog struct {
	Name   string
	DogAge int
}

// Years converts a human age into a dog age.
func Years(humanAge int) int {
	return humanAge * 7
}
