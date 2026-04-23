package main

import "fmt"

func arrayExample() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 10
	fmt.Println("dcl:", a)

	a[2] = 2 + 2
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])

	fmt.Println("length:", len(a))

	b := [3]int{20, 40, 99}
	fmt.Print(b)

	c := [...]int{100, 555, 30, 1002}
	fmt.Println(c)

	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2d: ", twoD)

	var twoD2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d2: ", twoD2)
}
