package main

import "fmt"

func Rotate(s []int, n int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	q := [...]int{1, 2, 3, 4, 5, 6}
	Rotate(q[:], 5)
	fmt.Println(q)
}