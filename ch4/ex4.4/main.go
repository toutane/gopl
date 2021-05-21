package main

import "fmt"

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)
	fmt.Println(a)
}

// Rotate rotate the slice of n elements.
func rotate(s []int, n int) {
	var buf []int
	for k := 0; k < n; k++ {
		buf = append(buf, s[k])
	}
	for i := 0; i < len(s)-n; i++ {
		s[i] = s[i+n]
	}
	for j, z := len(s)-n, 0; j < len(s); j, z = j+1, z+1 {
		s[j] = buf[z]
	}
}
