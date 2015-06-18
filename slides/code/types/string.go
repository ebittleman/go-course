// +build !appengine OMIT

package main

import "fmt"

func main() {
	StringExample()
}

func StringExample() {
	var stringExample string = "Here is a sentence"
	var formatExample string = "%s is number %d"

	fmt.Printf("Type: %T, Value: %v\n", stringExample, stringExample)

	fmt.Printf(formatExample, "Five", 5)
}

func PrintTypeAndValue(v interface{}) {

}
