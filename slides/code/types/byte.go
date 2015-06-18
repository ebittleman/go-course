// +build !appengine OMIT

package main

import "fmt"

func main() {
	ByteExample()
}

func ByteExample() {
	var byteExample1 byte = 0xff // base 16
	var byteExample2 byte = 0377 // base 8
	var byteExample3 byte = 255  // base 10

	PrintTypeAndValue(byteExample1)
	PrintTypeAndValue(byteExample2)
	PrintTypeAndValue(byteExample3)
}

func PrintTypeAndValue(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
