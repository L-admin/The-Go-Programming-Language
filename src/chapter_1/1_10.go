package main

import "fmt"

func reverser1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func resever2(s [3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	p := [...]int{1, 2, 3}
	reverser1(p[:])
	fmt.Println(p)

	q := [...]int{1, 2, 3}
	resever2(q)
	fmt.Println(q)
}

