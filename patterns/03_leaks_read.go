package main

import "fmt"

/*
1. Find the issue with below code. Understand the root cause of it.
2. Fix the issue with using select `cancellation` or `timeout`.
*/
func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("Done")
}
