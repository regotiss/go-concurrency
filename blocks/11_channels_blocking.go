package main

import "fmt"

/**
Channels in Go are said to be blocking.  this means that any goroutine that attempts to write to a channel that is full
will until the channel has ben emptied, and any goroutine that attempts to read from a channel that is empty will wait
until at least one item is placed on it.
*/
func main() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 { // ensure that we never write to the channel
			return
		}
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream) // block until something is put on the channel
}
