// practising slices

package main

import "fmt"

func main() {

	// declare array
	a := [5]int{0, 1, 2, 3, 4}
	
	// declare slice
	var s []int = a[1:3]
	
	// access slice + indices
	fmt.Println(s[1])
}
