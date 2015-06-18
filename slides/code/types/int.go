// +build !appengine OMIT

package main

import "fmt"

func main() {
	IntExample()
}

func IntExample() {
	var int8Max int8 = 127
	var int8Min int8 = -128

	// ~32 Thousand - Kilo
	var int16Max int16 = 32767
	var int16Min int16 = -32768

	// ~2 Billion - Giga
	var int32Max int32 = 2147483647
	var int32Min int32 = -2147483648

	// ~9 Quintillion - Exa
	var int64Max int64 = 9223372036854775807
	var int64Min int64 = -9223372036854775808

	PrintTypeAndValue(int8Max)
	PrintTypeAndValue(int8Min)
	PrintTypeAndValue(int16Max)
	PrintTypeAndValue(int16Min)
	PrintTypeAndValue(int32Max)
	PrintTypeAndValue(int32Min)
	PrintTypeAndValue(int64Max)
	PrintTypeAndValue(int64Min)
}

func PrintTypeAndValue(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
