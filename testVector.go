package main

import (
	"fmt"
)

func main() {
	var i int
	var q int

	print("enter the amount: ")
	fmt.Scan(&q)

	var v = make([]int, q)
	// You can create slices for your vector undefined 

	for i = 0; i < q; i++ {
		fmt.Print("enter with vector[", i, "]: ")
		fmt.Scan(&v[i])
	}

	fmt.Println(v)
}