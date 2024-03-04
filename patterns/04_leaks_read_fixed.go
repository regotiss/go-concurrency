package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		//doWork := func(done context.Context, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	//ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	terminated := doWork(done, nil)
	go func() {
		// Cancel the operation after 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}
