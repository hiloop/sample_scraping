package main

import (
	"fmt"
	"strconv"
)

func main2() {
	fmt.Printf("Hello World!\n")
	var start int = 19984
	var qty int = 10
	for i := start; i < start+qty; i++ {
		fmt.Println("fuck:" + strconv.Itoa(i))
	}
}
