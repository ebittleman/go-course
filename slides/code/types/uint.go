// +build !appengine OMIT

package main

import "fmt"

func main() {
	UintExample()
}

func UintExample() {
	var uint8Max uint8 = 255
	var uint8Min uint8 = 0

	var uint16Max uint16 = 65535 // ~65 Thousand - Kilo
	var uint16Min uint16 = 0

	var uint32Max uint32 = 4294967295 // ~4 Billion - Giga
	var uint32Min uint32 = 0

	var uint64Max uint64 = 18446744073709551615 // ~18 Quintillion - Exa
	var uint64Min uint64 = 0

	PrintTypeAndValue(uint8Max)
	PrintTypeAndValue(uint8Min)
	PrintTypeAndValue(uint16Max)
	PrintTypeAndValue(uint16Min)
	PrintTypeAndValue(uint32Max)
	PrintTypeAndValue(uint32Min)
	PrintTypeAndValue(uint64Max)
	PrintTypeAndValue(uint64Min)
}

func PrintTypeAndValue(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
