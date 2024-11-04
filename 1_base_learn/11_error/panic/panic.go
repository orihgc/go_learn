package main

import "fmt"

func main() {
	// 使用 defer + recover 来处理 panic 异常
	// 有点像finally 语句块
	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello, panic: %v\n", data)
		}
		fmt.Println("恢复之后从这里继续执行")
	}()

	panic("Boom")
	fmt.Println("这里将不会执行下来")
}
