package main

import (
	"fmt"
	"strings"
)

func strJoin(values...string) string{
	s := strings.Join(values, "")
	return s
}

func main(){
	fmt.Println(strJoin("hello", "world"))
}