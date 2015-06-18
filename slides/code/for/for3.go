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
		"Pluto",
	}

	for index := 0; index < len(planets); index++ {
		fmt.Println(index, planets[index])
	}
}
