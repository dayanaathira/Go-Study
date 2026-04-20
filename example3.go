package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func maths() {
	fmt.Println(s)

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Round(2.7))

	const a = 10
	const b = 20
	fmt.Println(a + b)

	const c = 30
	const d = 40
	fmt.Println(math.Floor(c / d))
}
