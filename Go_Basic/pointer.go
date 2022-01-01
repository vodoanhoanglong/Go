package main

import "fmt"

type Pointer struct {
	number int
}

func main() {
	var a int = 12
	var b *int = &a    // to address, var b *int = a -> error, &a -> get address of a
	fmt.Println(a, *b) // *b -> value, b -> address
	a = 24
	fmt.Println(a, *b)
	*b = 32
	fmt.Println(a, *b)

	// Pointer Arr -> slice
	pointerArr := []int{1, 2, 3}
	temp := pointerArr
	fmt.Println(temp)

	// Pointer Struct
	var pointerStruct *Pointer
	pointerStruct = &Pointer{number: 10} // or new(Pointer)
	fmt.Println(pointerStruct)

	pointerStruct.number = 20         // or (*pointerStruct).number = 20
	fmt.Println(pointerStruct.number) // or (*pointerStruct).number -> 20

	// Note: Map, when assigned to a new variable, points to the same address
	ex := map[string]string{"Long": "Yen", "Yen": "Long"}
	exCopy := ex
	exCopy["Long"] = "Love"
	fmt.Println(ex, exCopy) // => map[Long:Love Yen:Long] map[Long:Love Yen:Long]
}
