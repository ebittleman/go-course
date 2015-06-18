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

	for index, planetName := range planets {
		fmt.Println(index, planetName)
	}
}
