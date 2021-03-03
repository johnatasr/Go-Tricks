package main

import (
	"fmt"
)

type AudioOutput int

const (
    OutMute AudioOutput = iota // 0
    OutMono                    // 1
    OutStereo                  // 2
    _
    _
    OutSurround                // 5
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
