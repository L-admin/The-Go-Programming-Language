package main

import "fmt"

func isIsomorphism(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var charCntMap1 = map[uint8]int{}
	for i := 0; i != len(s1); i++ {
		charCntMap1[s1[i]]++
	}

	var charCntMap2 = map[uint8]int{}
	for i := 0; i != len(s2); i++ {
		charCntMap2[s2[i]]++
	}

	return eauql(charCntMap1, charCntMap2)
}

func eauql(x, y map[uint8](int)) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}


func main()  {
	s1 := "abc"
	s2 := "cba"
	s3 := ""
	s4 := "ghj"

	fmt.Println(isIsomorphism(s1, s2))
	fmt.Println(isIsomorphism(s1, s3))
	fmt.Println(isIsomorphism(s1, s4))
}
