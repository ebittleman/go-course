// +build !appengine OMIT

package main

import "fmt"

func main() {
	for index := 0; index < 4; index++ {
		fmt.Println(index)
	}
}
