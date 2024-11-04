package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// ChangeName 结构体接收器 尽量不要用
// 因为它会拷贝一份数据，然后在拷贝的数据上进行操作
// 如果内部的数据量比较大，那么这个拷贝的性能开销就会比较大
func (u User) ChangeName(newName string) {
	u.Name = newName // 无效
}

// ChangeAge 指针接收器，会改变原始数据
// 因为它会直接操作原始的数据
// 所以在内部的数据量比较大的时候，可以考虑用指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

// Handle TODO type a b 一般是下面这样写
type Handle func()

func (h Handle) Hello() {

}
