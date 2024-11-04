package main

import "fmt"

func main() {
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v,len: %d, cap: %d \n", a1, len(a1), cap(a1))

	var a2 [3]int
	fmt.Printf("a2: %v,len: %d, cap: %d \n", a2, len(a2), cap(a2))

}
