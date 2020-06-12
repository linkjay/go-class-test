# Go Classes Proof-of-Concept

This little project was created by me to propose a way of doing classes within Go. It features two basic classes `mather.go` and `printer.go` within the `lib` folder which can be easily imported and then initialized in code with a `lib.Mather` or `lib.Printer`.

It gets a bit more complicated when you want to introduce a constructor. That is possible, but it requires a function in the package to run the code and then return the class.

## Quirks

- Constructors don't get to be instantiated with `lib.CLASS` but have to have their `Init()` function called.
- The `randomizer.go` library will give the same string if the Unix timestamp is the same. This is intended, but is still a quirk because if the app were constantly running, it would be able to keep a randomized seed each time easily.

## Sample Output

```
$ go run .

What happens when you add 2 and 5 with our library is: 7
Letters being used: [A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z]
Seed that we started with: 1591950872
Random characters: GqEKDCJpbwgJllHjksRbZwVAuqxmOPywAJApGsKWyxYvOwlGiF
```