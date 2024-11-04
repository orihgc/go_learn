package main

import "fmt"

func main() {
	r1 := printNumWith2(3.1415926)
	println(r1)

	r2 := printBytes([]byte("Hello, Go!"))
	println(r2)
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return fmt.Sprintf("%02x", data)
}
