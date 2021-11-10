package main

import (
	"fmt"
	"strconv"
)

<<<<<<< HEAD
const Pi = 3.14
const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	/* Pi = 3
	fmt.Println(Pi) */

	fmt.Println(URL, SiteName)
	fmt.Printf("%T\n", URL)
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)
=======
func main() {
	/*
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		fl := 10.5
		i3 := int(fl)
		fmt.Printf("i3 = %T\n", i3)
		fmt.Println(i3)
	*/

	var s string = "100"
	fmt.Println(s)
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
>>>>>>> 27ecbf673ca7e8785721e5064e1ce87ecf556992

	fmt.Println(c0, c1, c2)
}
