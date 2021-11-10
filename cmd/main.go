package main

import "fmt"

func CallFunction(f func()) {
	f()
}

func Hello() {
	fmt.Println("Hello")
}

func main() {
	CallFunction(Hello)
}
