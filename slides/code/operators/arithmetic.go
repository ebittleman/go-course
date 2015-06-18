// +build !appengine OMIT

package main

import "fmt"

func main() {
	var onePlusOne int = 1 + 1
	var fourMinusThree int = 4 - 3
	var twoTimesTwo int = 2 * 2
	var sixDividedByTwo int = 6 / 2

	fmt.Println("1 + 1 = ", onePlusOne)      // addition
	fmt.Println("4 - 3 = ", fourMinusThree)  // subtraction
	fmt.Println("2 * 2 = ", twoTimesTwo)     // multiplication
	fmt.Println("6 / 2 = ", sixDividedByTwo) // division
}
