package main

import (
	"fmt"
	"sync"
	"time"
)

/**

WaitGroup allows you to wait for a set of concurrent operations to complete when you either
don't care about the result off the concurrent operation, or you have other means of collecting
their results.

If neither of those conditions are true, a `select` statement may be better suited.

It still doesn't guarantee order of completion.
**/

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Add a argument of 1 to indicate that one goroutine is beginning
	go func() {
		defer wg.Done() // Indicate to the WaitGroup we are done.
		fmt.Println("1st goroutine sleeping....")
		time.Sleep(1)
	}()

	wg.Add(1) // Add 1 to indicate goroutine is beginning
	go func() {
		defer wg.Done() // Indicate we are done
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // will block the main goroutine until all goroutines have exited
	fmt.Println("All goroutines complete")
}
