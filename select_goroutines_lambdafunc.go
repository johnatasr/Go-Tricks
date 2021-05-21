package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)
	quit := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "This one channel"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		two <- "This two channel"
	}()

	for {
		select {
		case rec1 := <-one:
			fmt.Println("Received from channel one", rec1)
		case rec2 := <-two:
			fmt.Println("Received from channel two", rec2)
		case <-quit:
			return
		}

	}

}
