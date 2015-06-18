// +build !appengine OMIT

package main

import "fmt"

func main() {
	var planets []string = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	for index := 0; index < 8; index++ {
		fmt.Println(index, planets[index])
	}
}
