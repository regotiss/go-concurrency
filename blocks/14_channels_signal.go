package main

import (
	"fmt"
	"sync"
)

/**
Closing a channel is also one of the ways you can signal multiple goroutines simultaneously.  Since closed channel can
read from an infinite number of times, it doesn't matter how many goroutines are waiting on it.

Remember `sync.Cond` type can perform the same behaviour. So why use channels, it's cheap and fast.  Channels are also
composable.
*/
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // Here the go routine waits until it is told it can continue
			fmt.Printf("%v has begin\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) // Here close the channel, thus unblocking all the goroutines.
	wg.Wait()
}
