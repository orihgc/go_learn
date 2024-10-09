package main

import "fmt"

// type a = b 是类型别名
func main() {
	var n News = fakeNews{
		Name: "hello",
	}
	n.Report()
}

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}

type fakeNews = News
