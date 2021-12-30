package main

import (
	"fmt"
	"strconv"
)

var Local string = "hello" // global variable by write first char uppercase

func main() {
	dec := 10.5                    // quick variable declare
	fmt.Printf("%v, %T", dec, dec) // %v is value %T is data type
	fmt.Println()
	var floatDec int = int(dec) // convert data type
	fmt.Printf("%v, %T", floatDec, floatDec)
	var strConv string = strconv.Itoa(42) // If don't use strconv the value is *, else the value is "42"
	fmt.Println()
	fmt.Printf("%v, %T", strConv, strConv)
}
