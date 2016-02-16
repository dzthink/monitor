package main

import (
	"fmt"
)

type Name struct {
	first string
	last string
}


func main() {
	list := map[string][]string{
		"a":{"1","2","3"},
		"b":{"10","20", "30"},
	}
	fmt.Println(list["a"])
}