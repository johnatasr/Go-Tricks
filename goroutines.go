package goroutines

import (
	"fmt"
	"time"
)

func sendInformation(channel chan<- string, info string) {
	channel <- info
}

func printInformation(channel <-chan string) {
	fmt.Println(<-channel)
}

func main() {

	// In async way
	information := make(chan string)

	go sendInformation(information, "Hello, it's a test")
	go printInformation(information)

	time.Sleep(5 * time.Second)

	//Can store mora than one in channel
	anotherinofrmation := make(chan int, 4)

	anotherinofrmation <- 2
	anotherinofrmation <- 3
	anotherinofrmation <- 4
	anotherinofrmation <- 5

	fmt.Println(<-anotherinofrmation) // 2
	fmt.Println(<-anotherinofrmation) // 3
	fmt.Println(<-anotherinofrmation) // 4
	fmt.Println(<-anotherinofrmation) // 5

}
