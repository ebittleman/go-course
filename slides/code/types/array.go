// +build !appengine OMIT

package main

import "fmt"

func main() {
	ArrayExample()
}

func ArrayExample() {
	var daysOfWeek []string = []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	PrintArray(daysOfWeek)

	fmt.Println("Day 1 is", daysOfWeek[1])
}

func PrintArray(array []string) {
	for i, item := range array {
		fmt.Println(i, item)
	}
}
