// +build !appengine OMIT

package main

import "fmt"

func main() {
	BoolExample()
}

func BoolExample() {
	var trueBool bool = true
	var falseBool bool = false

	PrintTypeAndValue(trueBool)
	PrintTypeAndValue(falseBool)
}

func PrintTypeAndValue(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
