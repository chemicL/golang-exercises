package main

import (
	"fmt"
	p "github.com/chemicL/golang-exercises/present"
)

func main() {
	arrays()
	slices()
	sliceAppends()
}

func arrays() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a // Creates an array of same size AND copies the contents.
	b[2] = 100

	fmt.Println(a, b)
}

func slices() {
	// Slices do not have fixed length.
	// Array's length is part of it's type, but with slices it's just a property of an instance.
	// Assigning a slice to another does not copy the contents. Slice holds a pointer to underlying array.

	a := []int{1, 2, 3, 4, 5}
	b := a[2:] // b simply starts at different offset using the same underlying array

	b[0] = 0
	fmt.Println(a, b)

	multByTwo(a)
	fmt.Println(a)
}

// The argument s is in fact a new slice instance, pointing to the same underlying array,
// therefore we can make changes without returning anything.
// In fact a slice is a struct type (just go to runtime/slice definition)
func multByTwo(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func sliceAppends() {
	p.Header("Basic append.")
	a := []int {1, 2, 3} // Note Go allocates a slice with capacity=10
	p.SliceInfo(a)

	// Append single item to a.
	a = append(a, 6)
	p.SliceInfo(a)

	// Why assign? From the documentation of append:
	//
	// The append built-in function appends elements to the end of a slice.
	// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
	// If it does not, a new underlying array will be allocated.
	//
	// Append returns the updated slice.
	//
	// It is therefore necessary to store the result of append, often in the variable holding the slice itself:
	// slice = append(slice, elem1, elem2)
	// slice = append(slice, anotherSlice...

	// Let's try something else.
	p.Header("Tricky append to a slice.")

	a = []int{}
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	p.SliceInfo(a)

	b := a[:1] // [1] with len = 1 cap = 3
	p.SliceInfo(b)

	c := append(b, 100) // Append in place of 2, b's len and cap don't change!
	p.SliceInfo(a)
	p.SliceInfo(b)
	p.SliceInfo(c)

	// Note: append modifies the underlying array, to grow arrays without modifying anyone else's slice view,
	// use copy.
	p.Header("Using copy to append")
	x := []int{}
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	p.SliceInfo(x)

	var y []int = make([]int, 1, cap(x)) // len = 1, cap = cap(x) as we want to grow to the same size as x
	copy(y, x) // [1] with len = 1 cap = 4
	p.SliceInfo(y)

	z := append(y, 100) // Append to y causes creating a new underlying array.
	p.SliceInfo(x)
	p.SliceInfo(y)
	p.SliceInfo(z)

	// Let's have a look at underlying arrays.
	zFull := z[:cap(z)]
	p.SliceInfo(zFull)
}

func double(s []int) {
	s = append(s, s...)
}