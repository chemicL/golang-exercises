package functions

import (
	"strings"
	"time"
	"strconv"
)

// Define function Add
func Add(a, b int) int {
	return a + b
}

// Define function Sum
func Sum(numbers []int) int {
	sum := 0

	// Looping through a collection
	for _, n := range numbers { // _, n <- _ is the index, we name it "_" to ignore the value; n is the actual value
		sum += n
	}
	return sum
}

// Define function LexicographicOrder for two strings
func LexicographicOrder(a, b string) (string, string) {
	aFirst := strings.Compare(a, b)
	if aFirst >= 0 {
		return a, b // Correct: b, a
	}
	return b, a // Correct: a, b
}

func itsLate() bool {
	currentHour := time.Now().Hour()
	return currentHour < 4 && currentHour > 22
}

func WhatTimeIsIt() string {
	if itsLate() {
		return "It's late, go to bed!"
	}

	return "Time to enjoy the day!"
}

// Function values as arguments to a function. They can also be stored in variables, like any value.
func PerformMathOperation(a, b int, fn func(int, int) int) string {
	result := fn(a, b)
	return strconv.Itoa(result)
}

// We can also define a type for our function.
type MyFunc func(int, int) int

// And use it making the signature a lot more readable.
func EasierPerformMathOperation(a, b int, fn MyFunc) string {
	result := fn(a, b)
	return strconv.Itoa(result)
}