/*
This is a block comment.
 */

// This is a line comment.

// Declare this package.
package main

// Import a package.
import "fmt"

// Declare a constant.
// Anything that starts with Uppercase letter is exported (variable at package level, function/method name/field name).
// Exporting means that another package can access it. All identifiers starting with lowercase letter are not exported.
const Name = "Alice"

// The entry point to application - function main in package main.
func main() {
	// Notice the indentation - Go enforces the use of Tab character for indentation.

	// ----------------------------------------
	// Print something to stdout.
	// ----------------------------------------

	fmt.Println("Hello, Go!")

	// ----------------------------------------
	// Declare a variable using zero value.
	// ----------------------------------------

	var aNumber int

	// Printing to stdout is performed using fmt package. Check the formatting directives at https://golang.org/pkg/fmt/.
	fmt.Println(fmt.Sprintf("The number is: %d", aNumber))

	// ----------------------------------------
	// Assign it.
	// ----------------------------------------

	aNumber = 5
	fmt.Println(fmt.Sprintf("Now the number is: %d", aNumber))

	// ----------------------------------------
	// Cast it to another type.
	// ----------------------------------------

	var floatNumber float32
	floatNumber = float32(aNumber) // No implicit casts.
	fmt.Println(fmt.Sprintf("Float number value %v", floatNumber))

	// ----------------------------------------
	// Declare and assign, with type inference.
	// ----------------------------------------

	var someWord = "hello"
	fmt.Println(fmt.Sprintf("The word given is: %s", someWord))

	// ----------------------------------------
	// Declare, assign, with type inference, without var keyword.
	// ----------------------------------------

	anotherWord := "Hi"
	fmt.Println(fmt.Sprintf("The other word chosen is: %s", anotherWord))
}
