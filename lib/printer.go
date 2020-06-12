package lib

import (
	"fmt"
	"os"
	"strconv"
)

// Printer is a print class that is capable of printing, and not printing! It also has a cool little function "IntTostring" to do those awesome integer -> string conversions!
type Printer struct{}

// Print is a printing function to print using fmt.println
func (Printer) Print(str string) {
	fmt.Println(str)
}

// Printf will print a formatted string!
func (p Printer) Printf(str string, formats ...interface{}) {
	p.Print(fmt.Sprintf(str, formats...))
}

// DontPrint is a function that just wont print because we needed testing material!!!
func (Printer) DontPrint() {
	os.Exit(1)
}

// IntToString is a function to convert ints into strings!
func (Printer) IntToString(num int) string {
	return strconv.Itoa(num)
}
