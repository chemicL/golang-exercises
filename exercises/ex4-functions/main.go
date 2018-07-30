package main

import (
	"github.com/chemicL/golang-exercises/present"
	"fmt"
	"github.com/chemicL/golang-exercises/exercises/ex4-functions/functions"
)

func main() {
	present.Header("Tests")
	fmt.Println("Check and run tests defined in functions_test.go instead of running this example.")

	present.Header("Btw, the time is:")
	fmt.Println(functions.WhatTimeIsIt())
}
