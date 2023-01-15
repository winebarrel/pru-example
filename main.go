package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

func main() {
	foo()
	fmt.Println("hello, world")
	bar()
}

func bar() {
	fmt.Println("bar")
}
