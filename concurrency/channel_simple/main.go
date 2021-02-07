package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	// Initializing the channel
	c := make(chan int)

	// Declaring and initilizing a RECEIVE-ONLY channel
	c1 := make(<- chan string)

	// Declaring and initilizing a SEND-ONLY channel
	c2 := make(chan <- string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2) // => chan int, <-chan string, chan<- string

	// Sending a value into the channel
	go f1(10, c)
	// Receiving a value from the channel
	n := <- c
	// Waiting for a value to be sent into the channel and print out that value
	fmt.Println("Value received:", n)
	fmt.Println("Exiting main ...")

    // Closing a channel
	close(c)
}
