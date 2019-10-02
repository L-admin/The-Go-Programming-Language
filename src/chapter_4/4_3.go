package main

import "fmt"

func Reverse3Array(s *[3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {
	q := [...]int{1, 2, 3}
	Reverse3Array(&q)
	fmt.Println(q)
}
