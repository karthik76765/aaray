package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("a")
	fmt.Println("a[4]")
}
