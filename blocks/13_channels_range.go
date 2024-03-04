package main

import "fmt"

/**
Since we can close a stream this opens up a few patterns.

First is ranging over a channel.  The `range` keyword, used in conjunction with the `for` statement, supports channels
as arguments and will automatically break the loop when the channel is closed.

The range does not return the second boolean value.  The specifics of handling a closed channel are managed for you to
keep the loop concise.
*/

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // Ensure we close the stream before we exit the goroutine
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream { // Here we range over the stream once closed.
		fmt.Printf("%v ", integer)
	}
	fmt.Println()
}
