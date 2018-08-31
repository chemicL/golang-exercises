package main

import (
	"math"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	regularErrors()

	//errorsPkgWay()

	//scanWithPkgErrors()
}

func regularErrors() {
	number := -1

	// Call function that can return an error. See the definition.
	squared, e := sqrt(number)

	// Idiomatic error handling check. First - check if something bad happened. Then handle it.
	if e != nil {
		// We can also check the error type.
		if squaringError, ok := e.(SquaringError); ok {
			log.Fatalf("got squaring error (%T): %+v", squaringError, squaringError)
		}

		log.Fatalf("got some problem (%T): %+v", e, e)
	}

	// Otherwise, continue processing with the assumption that everything is ok.
	fmt.Println(fmt.Sprintf("sqrt(%d) = %f", number, squared))
}

// error is in fact an interface. Therefore, we can define something that satisfies the interface.
type SquaringError struct {
	Msg string
}

// Method Error returning a string is what we must define.
func (e SquaringError) Error() string {
	return e.Msg
}

// Function that can return an error.
// Return a tuple of expected result and an error in case of issues. Always as last value.
func sqrt(number int) (float64, error) {
	if number < 0 {
		// Return zero-value for the domain return value and a filled-in error.
		// Should we log something here?
		// The Go way is - handle errors once. That means - either log it, handle and act accordingly, or return without action.
		return 0, SquaringError{"provided negative number"}
		//return 0, fmt.Errorf("provided negative number") // Defining our own type is not necessary, sometimes it's enough to create error on the spot.
	}
	// Return domain value and a nil error, which means the result is usable.
	return math.Sqrt(float64(number)), nil
}

// A more helpful way of dealing with errors involves the https://github.com/pkg/errors package.
func errorsPkgWay() {
	number := -1
	squared, err := sqrtWithPkgErrors(number)
	if err != nil {
		log.Fatalf("couldn't square number: %+v", err) // %+v to print stack trace together with message
	}
	fmt.Println(fmt.Sprintf("sqrt(%d) = %f", number, squared))
}

func sqrtWithPkgErrors(number int) (float64, error) {
	if number < 0 {
		// Handy creation of errors with a stacktrace. Check more functions in errors codebase.
		return 0, errors.New("provided negative number")
	}
	return math.Sqrt(float64(number)), nil
}

func scanWithPkgErrors() {
	var hello, world string

	if _, err := scanStrings(&hello, &world); err != nil {

		// Add message to the error
		err2 := errors.WithMessage(err,"handling error")

		fmt.Println(fmt.Sprintf("Got error: %+v", err2))
	}
}

func scanStrings(a, b *string) (int, error) {
	n, err := fmt.Sscan("Hello ", a, b)

	fmt.Println(fmt.Sprintf("Read %d tokens.", n))

	if err != nil {
		// call Wrap or Wrapf when dealing with library-provided error.
		return n, errors.Wrap(err, "scanning failed")
	}

	return n, nil
}