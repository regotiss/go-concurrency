package main

import (
	"fmt"
	"time"
)

/**
At time you may find yourself wanting to combine one or more `done` channels into a single done channel that closes
if any of its components channels close.
*/

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} { // takes a variadic slice of channels
		switch len(channels) {
		case 0: // since we use recursion, we must set up termination criteria, first one if empty
			return nil
		case 1: // 2nd termination criteria, contains one element
			return channels[0]

		}

		orDone := make(chan interface{})
		go func() { // main body of function and where recursion occurs.  Create a go routine so that we can wait for messages
			// without blocking
			defer close(orDone)
			switch len(channels) {
			case 2: // because of recursion, every recursive call will have at least 2 channels
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} { // create channels that will be close after a specific time
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))

}
