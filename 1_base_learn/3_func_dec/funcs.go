package main

// Fun0 单个返回值，不需要括号括起来
func Fun0(name string, age int) string {
	return name + " " + string(age)
}

// Fun1 多个返回值
func Fun1(age int, name string) (int, string) {
	return age, name
}

// Fun2 Fun 多个返回值，返回值有名字，可以直接return
func Fun2(a string, b string) (age int, name string) {
	age = 13
	name = ""
	return
}

// Fun3 多个参数具有相同类型放在一起，可以只写一次类型
func Fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 13
	e = 14
	g = "hh"
	return
}

// Fun4 可变参数，可变参数必须放在最后
func Fun4(a string, b int, names ...string) {

	// 使用 names 变量，names 是 []string 类型，可以看作切片
}

func main() {

	a := Fun0("abc", 13)
	println(a)

	b, _ := Fun1(1, "abc")
	println(b)
}
