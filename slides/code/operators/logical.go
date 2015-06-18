// +build !appengine OMIT

package main

import "fmt"

func main() {
	fmt.Println(1 == 1) // is 1 equal to 1?
	fmt.Println(2 != 2) // is 2 not equal 2?
	fmt.Println(3 > 2)  // is 3 less than 2?
	fmt.Println(9 < 4)  // is 9 less than 4?
}
