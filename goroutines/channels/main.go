package main

import "fmt"

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func main() {
	channelX := make(chan int)
	channelY := make(chan int)
	go odd(channelX)
	go even(channelY)
	fmt.Println(<-channelX)
	fmt.Println(<-channelY)
	fmt.Println(<-channelX)
	fmt.Println(<-channelY)
}
