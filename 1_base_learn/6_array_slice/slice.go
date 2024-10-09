package main

import "fmt"

// TODO 共享slice的底层数组
func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v,len: %d, cap: %d \n", s1, len(s1), cap(s1))

	// len = 元素个数，cap = 容量
	s2 := make([]int, 3, 4)
	fmt.Printf("s2: %v,len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 1) //没有超出容量，不会扩容
	fmt.Printf("s2: %v,len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 2) //超出容量，会扩容
	fmt.Printf("s2: %v,len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4)
	fmt.Printf("s3: %v,len: %d, cap: %d \n", s3, len(s3), cap(s3))

}
