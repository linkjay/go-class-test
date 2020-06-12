package lib

import (
	"math/rand"
	"time"
)

// Randomizer is a class that can give us some random stuff!
type Randomizer struct {
	seed    int64
	letters [52]string
}

// RandomizerInit is a function to create our seed and do some setup work for randomizing! This class demonstrates how we can start with some values and use them throughout a class
func RandomizerInit() Randomizer {
	//initialize our starting seed and our letter dictionary
	r := Randomizer{
		time.Now().Unix(),
		[52]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	}

	return r
}

// RandomLetters is a function that can give us any amount of random letters that we want
func (r Randomizer) RandomLetters(amount int) string {
	var finalString string = ""
	var arrayLen int = len(r.letters) - 1

	rand.Seed(r.UseSeed())

	for i := 0; i < amount; i++ {
		finalString += r.letters[rand.Intn(arrayLen)]
	}

	return finalString
}

// UseSeed will return the current seed and make a new seed
func (r Randomizer) UseSeed() int64 {
	// Store our old seed
	var oldSeed int64 = r.seed

	// Generate new one
	r.seed = time.Now().Unix()

	// Send it off!
	return oldSeed
}

// GetSeed will print the current seed used by the randomizer
func (r Randomizer) GetSeed() int64 {
	return r.seed
}

// GetLetters will show the 52 letters (A-z) that we use to generate our random strings!
func (r Randomizer) GetLetters() [52]string {
	return r.letters
}
