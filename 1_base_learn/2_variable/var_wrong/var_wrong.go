package main

// :=只能用于局部变量，不能用于全局变量
//a:=13

var a = "hello"

const External = "包外可访问"

func main() {
	a := "hh"
	println(a)
	println(External)
}
