package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	for {
		fmt.Printf("a1: %v,len: %d, cap: %d \n", a1, len(a1), cap(a1))
		break
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)
	}

	for i, value := range a1 {
		fmt.Printf("%d => %d\n", i, value)
	}
}
