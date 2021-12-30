package main

import "fmt"

/*
Primitive Data Types
1. Boolean
2. Numerics
	Integer
		signed int
		unsigned uint
	Float
	Complex
3. Text
	String
	Byte - UTF-8
	Rune - UTF-32
*/

func main() {
	check := 1 == 2
	var num1 int8 = 10 // 1010
	var num2 int8 = 3  // 0011
	fmt.Printf("%v, %T\n", check, check)

	result := num1 & num2 // 0010
	fmt.Printf("%v, %T\n", result, result)
	result = num1 | num2 // 1011
	fmt.Printf("%v, %T\n", result, result)
	result = num1 ^ num2 // 1001
	fmt.Printf("%v, %T\n", result, result)
	result = num1 &^ num2 // 0100
	fmt.Printf("%v, %T\n", result, result)

	// bit shift
	var bit int8 = 8                       // 2^3 -> 0000 0100
	fmt.Printf("%v, %T\n", bit<<3, bit<<3) // 2^3 * 2^3 = 2^6 -> 0010 0000
	fmt.Printf("%v, %T\n", bit>>3, bit>>3) // 2^3 / 2^3 = 2^0 -> 0000 0001

	// float
	numFloat := 3.14e7 // 3.14 ^ 7
	fmt.Printf("%v, %T\n", numFloat, numFloat)

	// complex
	var numComplex complex64 = 1 + 2i                          // or var numComplex complex64 = complex(1,2)
	fmt.Printf("%v, %T\n", numComplex, numComplex)             // (1+2i)
	fmt.Printf("%v, %T\n", real(numComplex), real(numComplex)) // 1
	fmt.Printf("%v, %T\n", imag(numComplex), imag(numComplex)) // 2

	// string
	var text1 string = "Hello GO"
	var text2 string = ", I'm Long"
	textConcat := text1 + text2
	fmt.Printf("%v, %T\n", textConcat, textConcat)

	// byte
	var byteArr []byte = []byte(textConcat)
	fmt.Printf("%v, %T\n", byteArr, byteArr)         // arr unicode
	fmt.Printf("%v, %T\n", string(byteArr), byteArr) // string

	// rune
	var runeSpecial rune = 'À'
	fmt.Printf("%v, %T\n", runeSpecial, runeSpecial)         // unicode
	fmt.Printf("%v, %T\n", string(runeSpecial), runeSpecial) // string
}
