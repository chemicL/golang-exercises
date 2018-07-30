package functions_test

import (
	// Testing package used for obvious reasons
	"testing"
	"github.com/chemicL/golang-exercises/exercises/ex4-functions"
	"math"
)

// Test definition for Add function with one sample
func TestAddSumsTwoNumbers(t *testing.T) {
	a := 100
	b := 99

	// Using function from our package:
	sum := functions.Add(a, b)

	// If statement
	if sum != 199 {
		// Log error and fail test
		t.Errorf("Summing two numbers failed. Sum function thinks %d + %d = %d", a, b, sum)
	}
}

// Test LexicographicOrder by using more samples
func TestLexicographicOrder(t *testing.T) {

	// Prepare struct to name our sample input
	type PairOfWords struct {
		word1   string
		word2   string
		ordered []string
	}

	// Parameterized tests look like this
	samples := []PairOfWords{
		{"hi", "hello", []string{"hello", "hi"}},
		{"twenty", "two", []string{"twenty", "two"}},
	}

	// Iterate through samples and run the same logic
	for _, pair := range samples {
		word1, word2 := functions.LexicographicOrder(pair.word1, pair.word2)
		if word1 != pair.ordered[0] || word2 != pair.ordered[1] {
			t.Errorf("Expected words in order %v. Given: [%s %s]", pair.ordered, word1, word2)
		}
	}
}

func TestPerformingOperationYieldsExpectedString(t *testing.T) {

	type Example struct {
		opName         string
		input1         int
		input2         int
		op             func (int, int) int
		expectedResult string
	}

	add := func (a, b int) int {
		return a + b
	}

	mult := func(a, b int) int {
		return a * b
	}

	pow := func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}

	examples := []Example {
		{"add", 10, 2, add, "12"},
		{"add",  2, 10, add, "12"},
		{"mult", 10, 2, mult, "20"},
		{"mult", 2, 10, mult, "20"},
		{"pow", 10, 2, pow,"100"},
		{"pow", 2, 10, pow, "1024"},
	}

	for _, e := range examples {
		if result := functions.PerformMathOperation(e.input1, e.input2, e.op); e.expectedResult != result {
			t.Errorf("Running %s for %d and %d returned %s instead of %s", e.opName, e.input1, e.input2, result, e.expectedResult)
		}
	}
}
