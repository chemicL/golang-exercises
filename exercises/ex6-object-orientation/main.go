package main

import (
	"fmt"
	"time"
)

func main() {
	// Define instance of a struct
	action := ShutdownAction{3} // Should it be pointer type? Come back here when asked to and add &.
	//action := ShutdownActionWithText{LastWords: "Byee!"} // Should it be pointer type? Come back here when asked to and add &.

	// Call a method on the struct. Soon it will become clear how this is possible.
	action.NoDelay() // This is the same as (&action).NoDelay().

	// Perform some action. We'll explain step by step how we got here. Read on (this file, top-to-bottom).
	//var a1 *PrintAction = nil


	//perform(a1)
	perform(action)

	//nilReceivers()
	//emptyInterface()
}

// Let's define two *Action types.
type PrintAction struct {
	text string
}

type ShutdownAction struct {
	delaySeconds int
}

type ShutdownActionWithText struct {
	ShutdownAction
	LastWords string
}

// This is a method definition. The receiver argument is specified before function name.
// We can only define methods for types from the same package.
func (p PrintAction) Execute() {
	fmt.Println(p.text)
}

// ShutdownAction can also be executed.
func (s ShutdownAction) Execute() {
	if s.delaySeconds > 0 {
		fmt.Println(fmt.Sprintf("Will shutdown in %d seconds.", s.delaySeconds))
		time.Sleep(time.Duration(s.delaySeconds) * time.Second)
	}
	fmt.Println("Shutting down NOW.")
	// os.Exit(0) // Commented out for now ;-)
}

// Define method on pointer receiver.
// Reasons? Modifying the object (passing ownership) or performance (?).
// (?): Check why it's more expensive to use pointers: https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
// Rather focus on the ownership aspect.
//
// When specifying the methods prefer consistent API (always by pointer or always by value).
func (s *ShutdownAction) NoDelay() { // Try removing the * and see this method has no effect on the delay.
	s.delaySeconds = 0
}

// Notice a pattern here? It seems we can use the above in the same style.
// At point of use, we can define an interface that will depict this.
// The interface is implemented implicitly - there is no "implements" keyword.

type Action interface {
	Execute()
}

// Let's make Actions describe themselves

type Describable interface {
	Describe()
}

func (s *ShutdownAction) Describe() {
	describe(s)
}

func (p *PrintAction) Describe() {
	describe(p)
}

func describe(describable Describable) {
	// We can print the type of the receiver with %T
	fmt.Println(fmt.Sprintf("Describing receiver: (%+v, %T)", describable, describable))
}

func perform(action Action) {

	// Type assertion (can also use in switch statements).
	if d, ok := action.(Describable); ok { // Here is why mixing pointer and value receiver methods is not a good idea.
	// Action.Execute method is defined on value receivers, so we pass a value Action here.
	// Try changing the declaration of action in main() to pointer type and see that it works.
		d.Describe()
	} else {
		fmt.Println(fmt.Sprintf("Can't describe given action: %T", action))
	}

	fmt.Printf("Pointer to action: %p\n", action)

	// Check nullness
	if action == nil {
		fmt.Println("Trying to perform action on nil receiver")
		return
	}
	action.Execute()
}

func nilReceivers() {
	// We can also call methods on nil receiver if it's type is known.
	var printAction *PrintAction
	var nilAction Describable = printAction

	// interface value holds a nil concrete value, which is non-nil; try replacing printAction with nil.

	// What happens here? Under the covers, interface value is a tuple (value, type),
	// that's why having nil value but a defined type we can call methods on nil receiver.

	nilAction.Describe()
}

func emptyInterface() {
	// An interface can specify zero methods - it's called an empty interface. It can represent values of any type.
	describeAny(nil)
	describeAny(ShutdownAction{})
	var a *ShutdownAction
	describeAny(a)
}

// Redefine describe function which works on any value.
func describeAny(i interface{}) {
	fmt.Println(fmt.Sprintf("Got (%+v, %T)", i, i))
}