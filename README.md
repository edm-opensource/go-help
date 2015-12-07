# go-help
[![travis build](https://img.shields.io/travis/edm-opensource/go-help.svg)](https://travis-ci.org/edm-opensource/go-help)
[![codecov coverage](https://img.shields.io/codecov/c/github/edm-opensource/go-help.svg)](https://codecov.io/github/edm-opensource/go-help)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/edm-opensource/go-help/master/LICENSE)

A go helper library

Provides simple utilities that does not exist in the standard go libraries.


# Installation
To get the package, run:

    go get github.com/edm-opensource/go-help/...

# Usage
To use the go-help package just import it in your go file:

    import "github.com/edm-opensource/go-help/gohelp"

# Features

### Integers
* **ModPow:** Calculates power of big numbers modulus some other number

### Slices
* **RemoveDuplicateInts:** Removes all duplicates from an int slice
* **RemoveDuplicateStrings:** Removes all duplicates from a string slice
* **RemoveDuplicateFloat64:** Removes all duplicates from a float64 slice
* **ReverseInts:** Reverses a slice of ints
* **ReverseStrings:** Reverses a slice of strings
* **ReverseFloat64s:** Reverses a slice of float64s
* **RemoveAllInstancesOfStrings:** Returns a copy of the slice with all instance of the strings from the arguments removed
* **RemoveAllInstancesOfInts:** Returns a copy of the slice with all instance of the ints from the arguments removed
* **RemoveAllInstancesOfFloat64:** Returns a copy of the slice with all instance of the float64 from the arguments removed
* **ContainsString:** Returns true if slice contains input string
* **ContainsInt:** Returns true if slice contains input int
* **ContainsFloat64:** Returns true if slice contains input float64

### Strings
* **EncodeRomanNumerals:** Encodes an int to a string of roman numerals
* **DecodeRomanNumerals:** Decodes a string of roman numerals into an int
* **TitleFirstWord:** Changes the first letter of the first word to uppercase

# Contributing
Things we want you to contribute with:
* **Utilities you think are missing from the standard libraries or this package**
* **Issues regarding bugs or problems within out code**
* **Solutions to existing issues**
* **Feature requests**

### To contribute with code:
1. Fork this repository
2. Write code
3. Write tests, please
4. Confirm tests pass
5. Commit changes
6. Push to your repository
7. Make pull request to our develop-branch

Use camelCase for variable names

## File structure
The file structure is straight forward.
All go-files are in the go-help folder.
Each method is placed in the appropriate file, ie. methods concerning slices are in slices.go
Each file has a corresponding test-file where all methods must be tested.

### Writing tests
Please follow our way of writing tests in the appropriate files
``` go
func TestReverseStrings_1(t *testing.T) {
	input := []string{"one", "two", "three"}
	expected := []string{"three", "two", "one"}
	actual := ReverseStrings(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(formatError(input, expected, actual))
	}
}
```

The method formatError should always be used to print errors to ensure good readability in the console.

### Running tets
If you are using a Mac like us, you can add this method to your .bash_profile file:

``` bash
got(){
go test github.com/edm-opensource/go-help/... 
| sed ''/Expected/s//$(printf "\033[32mExpected\033[0m")/'' 
| sed ''/Actual/s//$(printf "\033[31mActual\033[0m")/'' 
| sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' 
| sed ''/ok/s//$(printf "\033[32mok\033[0m")/'' 
| sed ''/PASS/s//$(printf "\033[32mPass\033[0m")/''
} 
```

This will give you colored and easy-to-read reports using the command ``` got() ``` in your terminal

All contributions are welcome!

# License
MIT
