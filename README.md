# go-help
[![travis build](https://img.shields.io/travis/edm-opensource/go-help.svg)](https://travis-ci.org/edm-opensource/go-help)
[![codecov coverage](https://img.shields.io/codecov/c/github/edm-opensource/go-help.svg)](https://codecov.io/github/edm-opensource/go-help)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/edm-opensource/go-help/master/LICENSE)

A go helper library

Provides simple utilities that does not exist in the standard go libraries.


# Installation
To get the package, run:

    go get github.com/edm-opensource/go-help
    
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
We encourage you to contribute with utilities that you think are missing from the standard libraries.

Feel free to make pull requests, we will look at them as soon as possible.

https://codecov.io/github/edm-opensource/go-help


# License
MIT
