package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "Hello"
	// var s int = 100
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Printf("err = %T\n", err)
		fmt.Println(err)
	} else {
		fmt.Printf("i = %T\n", i)
	}
}
