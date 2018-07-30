package present

import "fmt"

func SliceInfo(s []int) {
	fmt.Println(fmt.Sprintf("len=%d cap=%d %v", len(s), cap(s), s))
}

func Header(msg string) {
	fmt.Println("----------------------------------------")
	fmt.Println(msg)
	fmt.Println("----------------------------------------")
}