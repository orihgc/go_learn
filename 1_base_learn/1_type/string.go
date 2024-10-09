package main

import "unicode/utf8"

func main() {
	print("Hello,Go\n")
	println("Hello,Go")

	println(len("你好"))
	println(utf8.RuneCountInString("你好abc"))

}
