package main

import (
	"fmt"
	// The following is a type alias. present package can be used via p.
	p "github.com/chemicL/golang-exercises/present"
	"os"
)

func main() {

	pointers()

	structs()

	arrays()

	slices()

	moreOnSlices()

	dynamicArraysUsingSlices()

	maps()
}

func pointers() {
	p.Header("Pointers.")

	someone := "Point At Me"
	// Declare a pointer.
	var p *string
	// Generate a pointer using &.
	p = &someone
	fmt.Println(p)
	// Dereference a pointer.
	fmt.Println(*p)
	// Assign the underlying value.
	*p = "I'm someone else"
	fmt.Println(someone)
}

// ---------------------------------------- 
// Structs.
// ----------------------------------------

// Define a struct type.
type Person struct {
	Name string
	Surname string
	Age int
}

func structs() {
	p.Header("Structs.")

	// Create instance of Person.
	walt := Person{"Walt", "Disney", 65}
	fmt.Println(walt)

	// Access fields.
	walt.Name = "Walter"

	fmt.Println(walt)

	// Pointer to struct.
	p := &walt

	p.Name = "Walty" // Accessing fields looks the same.
	fmt.Println(walt)

	// Create Person in multiple ways.
	var (
		p1 = Person{} // Name and Surname are empty strings, Age is zero.
		p2 = Person{Surname: "Doe", Name: "John"} // Age is zero. Order doesn't matter.
	)

	fmt.Println(p1, p2)
}

// ----------------------------------------
// Arrays and slices.
// ----------------------------------------

func arrays() {
	p.Header("Arrays.")

	// Declare variable hosts as a 2-element array. The length is part of its' type. That's why they can't be resized.
	var hosts [2]string
	hosts[0] = "mordor"
	hosts[1] = "shire"
	fmt.Println(hosts[0], hosts[1])

	heroes := [2]string{"Frodo", "Sam"}
	fmt.Println(heroes)

	// Iterate over an array - note that we get only the index.
	for i := range heroes {
		fmt.Println(i, heroes[i])
	}

	// Iterate over an array with index and value. We could ignore the index using _ identifier.
	for i, hero := range heroes {
		fmt.Println(i, hero)
	}
}

func slices() {
	p.Header("Slices.")

	// First, we create an array.
	a := [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Define a slice.
	// Slice doesn't contain any data, it describes a section of underlying array.
	var s = a[1:4] // [low, high) -> a[1] a[2] a[3]
	fmt.Println(s)

	s[0] = 100 // Reference 0 element of slice -> a[1] (slice defined as a[1:4])
	fmt.Println(a)

	// Slice literal - creates an array and a slice that references it.
	goodGuys := []string {"Anakin", "Luke", "Yoda"}

	fmt.Println(fmt.Sprintf("Good guys: %v", goodGuys[1:]))

	// Try defining a slice for type Person.
}

func moreOnSlices() {
	p.Header("More on slices.")

	s := []int {100, 200, 400, 800, 1600}
	p.SliceInfo(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	p.SliceInfo(s)

	// Extend its length.
	s = s[:4]
	p.SliceInfo(s)

	// Drop its first two values.
	// What happens to the first two items? Will they get GC'd?
	// No - unless you append to the slice and a new array is allocated.
	s = s[2:]
	p.SliceInfo(s)

	// Check what happens if we try to extend length beyond capacity:
	// s = s[:4]
	// p.SliceInfo(s)

	// Nil slice. No underlying array.
	s = nil
	p.SliceInfo(s)
}

func dynamicArraysUsingSlices() {
	p.Header("Dynamic arrays.")

	// We create a slice with given length or length and capacity using built-in make function:
	// https://golang.org/pkg/builtin/#make

	// Note on builtin - it's not a real package, it's a namespace for language's special identifiers.
	// E.g. make for slices can take three arguments, but the Godoc signature only shows one.

	uncapped := make([]int, 3) // 3 element underlying array
	p.SliceInfo(uncapped)

	capped := make([]int, 0, 5) // 0 length with capacity of 5
	p.SliceInfo(capped)

	// p.SliceInfo(capped[3:]) // [3:len(slice)] == [3:0]
	// p.SliceInfo(capped[3:4]) ?
	// p.SliceInfo(capped[3:3]) ?

}

func maps() {
	p.Header("Maps.")

	// Maps have keys which can be compared using ==
	// Declare variable for a mapping of string to string
	var m map[string]string // nil is the zero value

	fmt.Println(m)

	//m["Key"] = "Value" // won't work

	// We also use make to initialize a map
	m = make(map[string]string) // initialized map

	m["Key"] = "Value"

	fmt.Println(m)

	visited := map[string]string {
		"Poland": "Warsaw",
		"Spain": "Madrid",
		"United Kingdom": "London",
		"Belgium": "Brussels",
		"Netherlands": "Amsterdam",
		"France": "Paris",
		"Czech Republic": "Prague",
		"Austria": "Vienna",
		"Sweden": "Stockholm",
		"Hungary": "Budapest",
		"Portugal": "Lisbon",
		"Tunisia": "Tunis",
	}

	// We can provide initial capacity hint.
	// It doesn't mean any limit, maps grow to accomodate all items stored in them.
	// The actual size and growing policy is implementation-dependent.
	// An exception is a nil map, which acts like an empty map, but no items can be added to it.
	toVisit := make(map[string]string, 4)

	fmt.Println(fmt.Sprintf("The entire map: %+v.", visited))

	var capitalOfPoland = visited["Poland"]

	if capitalOfPoland != "Warsaw" {
		fmt.Print(fmt.Sprintf("Capital city of Poland is not Warsaw?! Got %s", capitalOfPoland))
		os.Exit(1)
	}

	// Iterate over the map (order is not specified) using both key and value.
	for key, value := range visited {
		fmt.Println(fmt.Sprintf("The capital of %s is %s.", key, value))
	}

	if _, haveVisited := visited["Italy"]; !haveVisited {
		toVisit["Italy"] = "Rome"
	}

	// Hm. I'm not sure whether I visited Tunis, I'll visit in the future just to be sure.
	delete(visited, "Tunisia")
	toVisit["Tunisia"] = "Tunis"

	// Iterate over the keys of the map.
	for country := range toVisit {
		fmt.Println("Will visit the capital of", country)
	}

	// Quick implementation of a concept known as map.values().
	var capitalsToVisit []string
	for _, city := range toVisit {
		capitalsToVisit = append(capitalsToVisit, city)
	}

	fmt.Println(fmt.Sprintf("I'll visit the following cities: %v", capitalsToVisit))
}
