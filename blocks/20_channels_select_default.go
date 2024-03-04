package main

import (
	"fmt"
	"time"
)

/**
Like case statements, the select statement also allows for a default clause in case you'd like to do something if all
channels you're selecting against are blocking.

Usually you'll see the default clause used in conjunction with a for-select loop.  This allows a goroutine to make
progress on work while waiting for another goroutine to report result.
*/

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Acheived %v cycles of work before signalled to stop.\n", workCounter)
}
