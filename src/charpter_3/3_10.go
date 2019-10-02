package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	if len(s) < 3 {
		return s
	}

	var buf bytes.Buffer
	for i := 0; i != len(s); i++ {
		if i != 0 && i % 3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main()  {
	s := ""
	s1 := "1"
	s2 := "123"
	s3 := "1234"
	s4 := "12345"
	s5 := "123456"

	fmt.Println(comma(s))
	fmt.Println(comma(s1))
	fmt.Println(comma(s2))
	fmt.Println(comma(s3))
	fmt.Println(comma(s4))
	fmt.Println(comma(s5))
}
