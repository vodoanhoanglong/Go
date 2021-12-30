package main

import "fmt"

func main() {
	// var autoPoints [3]int -> [0, 0, 0 ]
	points := [...]int{7, 10, 5, 2, 5}
	fmt.Println(len(points)) // get length arr

	pointer := &points // access same a memory
	// pointer := points // access diff a memory
	pointer[2] = 0

	fmt.Printf("%v, %T\n", points, points)   // ele 2 -> 0
	fmt.Printf("%v, %T\n", pointer, pointer) // ele 2 -> 0

	// slice
	slice := []int{7, 10, 5, 2, 5}
	sliceCopy := slice
	sliceCopy[2] = 1

	fmt.Printf("%v, %T\n", slice, slice)         // ele 2 -> 1
	fmt.Printf("%v, %T\n", sliceCopy, sliceCopy) // ele 2 -> 1

	// 1 2 3 4 5 6 7 8 9
	//         ^   ^
	// len is 3 : 5, 6, 7
	// cap is 5 : 5, 6, 7, 8, 9

	a := []int{2, 5, 10, 12, 45, 50}
	b := a[:]                                     // copy all
	c := a[3:]                                    // start index 3
	d := a[:5]                                    // start 0 to index < 5
	e := a[3:5]                                   // start 3 to index < 5
	e[1] = 100                                    // change all value of arr at that index
	fmt.Printf("%v, %v, %v\n", a, len(a), cap(a)) // [2 5 10 12 45 50], 6, 6
	fmt.Printf("%v, %v, %v\n", b, len(b), cap(b)) // [2 5 10 12 45 50], 6, 6
	fmt.Printf("%v, %v, %v\n", c, len(c), cap(c)) // [12 45 50], 3, 3
	fmt.Printf("%v, %v, %v\n", d, len(d), cap(d)) // [2 5 10 12 45], 5, 6
	fmt.Printf("%v, %v, %v\n", e, len(e), cap(e)) // [12 45], 2, 3

	// make
	makeArr := make([]int, 10) // make(arr, length, capacity)
	// [0 0 0 0 0 0 0 0 0 0], 10, 10
	fmt.Printf("%v, %v, %v\n", makeArr, len(makeArr), cap(makeArr))

	// append
	appendArr := make([]int, 0) // [], 0, 0
	/* appendArr = append(appendArr, 1) // append: insert ele => [1] , 1, 1 */
	/* appendArr = append(appendArr, 1, 2, 3, 4, 5) // insert multiple value */
	appendArr = append(appendArr, []int{1, 2, 3, 4, 5, 6}...) // insert a slice
	fmt.Printf("%v, %v, %v\n", appendArr, len(appendArr), cap(appendArr))

	// slice instance stack
	stack := []int{1, 2, 3, 4, 5, 6} // pop
	stack = stack[:5]
	fmt.Println(stack)

	// slice instance queue
	queue := []int{1, 2, 3, 4, 5, 6}
	queue = queue[1:] // pop
	fmt.Println(queue)

	// slice middle
	middle := []int{1, 2, 3, 4, 5, 6}
	middle = append(middle[:2], middle[3:]...)
	fmt.Println(middle)

}
