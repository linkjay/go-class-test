package main

import "web/lib"

func main() {
	//
	// We can import our library from the ./lib folder with the "web/lib" import above, and bring in the relating files to be used here
	//

	// our custom printing library that we use throughout this project for conversions and printing
	var print lib.Printer

	// custom math library to do some adding below!
	var math lib.Mather

	// let's establish two ints that we want to add
	var x, y = 2, 5

	// let's use our math lib to add our two ints
	var addedInt int = math.Add(x, y)

	// convert it to a string using our printing library!
	var finalAdd string = print.IntToString(addedInt)

	// print it using our printing lib!
	print.Printf("What happens when you add %d and %d with our library is: %s", x, y, finalAdd)

	// This class has a "constructor". There is a quirk with it which is that it has to be a static function that is global. The two classes above could be initialized here by doing lib.CLASS, but in order to execute code before hand, we actually have to call what is essentially now a global function which returns the class type. See "lib/randomizer.go" to see how this works
	var random = lib.RandomizerInit()

	// print out our letters with our specialized formatted printing function from our print library
	print.Printf("Letters being used: %v", random.GetLetters())
	print.Printf("Seed that we started with: %d", random.GetSeed())
	print.Printf("Random characters: %s", random.RandomLetters(50))
}
