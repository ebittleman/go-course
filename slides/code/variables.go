// +build !appengine OMIT

package main

import "fmt"

func main() {

	// declared but and automatically initialized to 0
	var someNumber int

	// declared and initialized
	var someOtherNumber int = 5

	// implicitly declared and initialized
	yetAnotherNumber := 10

	// lets see what the values are
	fmt.Println(someNumber, someOtherNumber, yetAnotherNumber)
}
