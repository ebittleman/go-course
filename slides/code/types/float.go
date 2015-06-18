// +build !appengine OMIT

package main

import "fmt"

func main() {
	FloatExample()
}

func FloatExample() {
	// Precision up to 8 Significant Figures
	var float32Dec float32 = 0.123456789123456789
	var float32Sci float32 = 9223372036854775807

	// Precsion up to 17 Significant Figures
	var float64Dec float64 = 0.123456789123456789123456789
	var float64Sci float64 = 9223372036854775807

	PrintTypeAndValue(float32Dec)
	PrintTypeAndValue(float32Sci)
	PrintTypeAndValue(float64Dec)
	PrintTypeAndValue(float64Sci)
}

func PrintTypeAndValue(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
