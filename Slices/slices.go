// practising slices

package main

import "fmt"

func main() {

	a := [5]int{0, 1, 2, 3, 4}

	var s []int = a[1:3]

	fmt.Println(s[1])
}
