package lib

// Mather is a library that can do basic math things. It's a good test of how we can implement multiple functions within the context of a file like this; it's like a class!
type Mather struct{}

// Add will go ahead and add two ints together
func (Mather) Add(x, y int) int {
	return x + y
}

// Sub will go ahead and substract two ints from each other
func (Mather) Sub(x, y int) int {
	return x - y
}
