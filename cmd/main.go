package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = "A"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 2
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
