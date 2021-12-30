package main

import (
	"fmt"
)

const constant = 42

func main() {
	fmt.Printf("%v, %T\n", constant, constant) // 42
	const constant = 12                        // shadow of constant variable in scope
	fmt.Printf("%v, %T\n", constant, constant) // 12

	// enum
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", blue, blue)
}

const (
	_ = iota // iota + index: init indentity start 0 + index
	// char _ is don't care, use it -> avoid wasting memory
	red    // 1
	yellow // 2
	blue   // 3
	Black  // 4 and this is a global variable
)
