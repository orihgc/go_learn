package main

import "fmt"

func main() {
	name := "张三"
	age := 18
	str := fmt.Sprintf("%s 今年 %d 岁", name, age)
	println(str)
}
