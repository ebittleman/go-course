// +build !appengine OMIT

package main

import "fmt"

func main() {
	MapExample()
}

func MapExample() {
	var daysOfWeek map[string]int = map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
	}

	PrintArray(daysOfWeek)

	fmt.Println("Wednesday is Day Number", daysOfWeek["Wednesday"])
}

func PrintArray(data map[string]int) {
	for i, item := range data {
		fmt.Println(i, item)
	}
}
