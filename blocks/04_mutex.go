package main

import (
	"fmt"
	"sync"
)

/**
Mutex stands for "mutual exclusion" and is a way to guard critical sections of your program.  A critical section is an
area of your program that requires exclusive access to a shared resource.

A Mutex provides a concurrent-safe way to express exclusive access to a shared resource.

Mutex share memory by creating a convention developers must follow to synchronize access to the memory.  You are
responsible for coordinating access to this memory by guarding access to it with a mutex.

*/

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()         // here we request exclusive use of the critical section, i.e count.
		defer lock.Unlock() // Here we indicate we are done with the critical section local is guarding
		// To use defer is a common idiom when using a Mutex to ensure the call always happens
		// even when a panic occurs.  Failing do so will result in a deadlock.
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()         // request lock
		defer lock.Unlock() // release lock
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	//Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	//Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
