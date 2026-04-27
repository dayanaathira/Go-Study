package main

import (
	"fmt"
)

func addNumber(a int, b int) int {
	return a + b
}

func multipleNumber(e, c int, d int) int {
	return c*d + e
}

func exampleFunc() {
	j := multipleNumber(1, 2, 3)
	k := addNumber(12, 34)

	fmt.Println(j)
	fmt.Println(k)
}
