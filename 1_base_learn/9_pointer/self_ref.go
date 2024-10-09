package main

func main() {

}

type Node struct {
	//自引用只能使用指针 编译的时候要确定结构体的大小，嵌套会导致结构体大小不确定
	//left Node
	//right Node

	// 指针的大小是固定的
	left  *Node
	right *Node

	// 这个也会报错 嵌套会导致结构体大小不确定
	// nn NodeNode
}

type NodeNode struct {
	node Node
}
