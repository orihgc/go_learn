package main

func main() {
	for i := 0; i < 10; i++ {
		println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}
