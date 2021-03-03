package main

import (
	"fmt"
)

const (
	a = iota + 1  // 0
	_             // 1 
	c             // 2
	d             // 3
)

func main() {		
	fmt.Println(a, c, d)  // 1 , 3, 4

}
