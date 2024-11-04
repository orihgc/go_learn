package main

// Global首字母大写，全局可访问
var Global = "全局变量"

// 首字母小写，只能在包内访问，子包也无法访问
var local = "局部变量"

var (
	First  string = "abc"
	second string = "def"
)

func main() {
	var a = 13
	println(a)

	var c uint = 12
	println(c)

	// go是强类型语言，类型不同，无法比较
	//println(a==c)

	// 只声明不赋值, 默认值为0
	var d int
	println(d)

}
