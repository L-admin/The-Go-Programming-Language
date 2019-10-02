package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, tmp := range os.Args[:]{
		s += sep + tmp
		sep = " "
	}
	fmt.Println(s)
}


