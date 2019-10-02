package main

import (
	"fmt"
	"sort"
)

func max(values...int) int{
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	return values[len(values)-1]
}

func min(values...int) int{
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	return values[0]
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}

