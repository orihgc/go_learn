package main

const inernal = "包内可访问"
const External = "包外可访问"

func main() {
	println(inernal)
	println(External)
}
