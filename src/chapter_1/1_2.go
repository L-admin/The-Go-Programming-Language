package main

import (
"fmt"
"os"
)

func main() {
	for i, tmp := range os.Args[:]{
		fmt.Println(i, tmp)
	}
}
