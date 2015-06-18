// +build !appengine OMIT

package main

import "fmt"

func main() {
	var ok bool = false

	if ok {
		fmt.Println("This statement is true")
	} else {
		fmt.Println("This statement is false")
	}
}
