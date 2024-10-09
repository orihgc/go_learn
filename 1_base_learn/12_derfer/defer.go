package main

import "fmt"

func main() {
	// 多个defer 函数，会按照先进后出的顺序执行
	// 在return 之前执行，在panic 之后执行
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
}
