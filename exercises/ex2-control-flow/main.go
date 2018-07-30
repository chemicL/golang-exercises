package main

import (
	"time"
	"fmt"
	"os"
	"github.com/chemicL/golang-exercises/present"
)

func main() {

	simpleControlFlow()
	
	switchWithoutCondition()

	deferredGoToSleep()

}

func simpleControlFlow() {
	present.Header("Simple control flow.")

	// ----------------------------------------
	// For loop; No parentheses.
	// ----------------------------------------
	
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(fmt.Sprintf("Sum of numbers 1 to 9 is %d", sum))

	// ----------------------------------------
	// While loop (for's init and post statements are optional).
	// ----------------------------------------
	
	count := 1
	for count < 10 {
		count++
	}
	fmt.Println("Done counting to 10.")

	// ----------------------------------------
	// Loop forever (commented out, yes...).
	// ----------------------------------------
	
	// for {
	//
	// }

	// ----------------------------------------
	// If statement.
	// ----------------------------------------
	
	if count > sum { // { } Braces always required
		fmt.Println("Oh. Something is completely wrong. 10 is less than the sum of numbers 0..10")
		os.Exit(1)
	} else { // typically in the above case (returning/exiting) we'd omit else and assume any code that follows
		// is executed if the check didn't catch it. That's idiomatic Go's style.
		fmt.Println("We are all good.")
	}

	// ----------------------------------------
	// Enhanced if - with init.
	// ----------------------------------------
	if word := "hello"; len(word) < 4 {
		os.Exit(1)
	}

	// ----------------------------------------
	// Switch statement; Evaluated from top to bottom.
	// ----------------------------------------
	
	switch firstLetter := "hello"[0]; firstLetter {
	case 'h':
		fmt.Println("Yep. We're all good here.")
		// No break statement! Go just runs the selection.
	default:
		fmt.Println("The World is about to end.")
		os.Exit(1)
	}
}

func switchWithoutCondition() {
	present.Header("Switch without condition.")

	t := time.Now()
	switch {
	case t.Hour() < 2:
		fmt.Println("Go to bed, NOW!")
	case t.Hour() < 6:
		fmt.Println("Shhhh..")
	case t.Hour() < 8:
		fmt.Println("Time to go to work!")
	case t.Hour() < 17:
		fmt.Println("Work hard!")
	case t.Hour() < 23:
		fmt.Println("Enjoy the evening.")
	}
}

func deferredGoToSleep() {
	present.Header("Deferred go to sleep.")

	defer fmt.Println("Fell asleep.")

	fmt.Println("Can't sleep... Will count sheep.")

	for i := 0; i < 10; i++ {
		/*defer */fmt.Println(fmt.Sprintf("Sheep %d", i))
	}

	// Deferred code executes.
	// Try counting from the end using defer.
}
