package main

import (
	"fmt"
	"time"
)

func main() {

	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()

	fmt.Println(ReturnClosure("Tom")())

	Delay()
	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	// 这里的i 是引用，不是值
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	// 这里会打印 10 个 10
	// 因为for 循环结束，i 的值是10
	// 所以所有的匿名函数都会引用 i 的最后一个值
	for _, fn := range fns {
		fn()
	}
}
