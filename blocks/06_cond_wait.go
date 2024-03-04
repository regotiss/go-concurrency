package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 Cond type is a rendezvous point for goroutine waiting for or announcing the occurrence of an event.

In that definition, an "event" is any arbitrary signal between two or more goroutines that carries no information
other than the fact that it has occurred.  Very often you'll want to wait for one of these signals before continuing
execution on a goroutine.

Naive approach would be to use an infinite loop

for conditionTrue() == false {
}

This would consume all the cycles of one core.  We could add a sleep to fix it.

for conditionTrue() == false {
  time.Sleep(1*time.Millisecond)
}

Still inefficient, since you need to figure out how long to sleep for.  Would be better if we could let a goroutine sleep
until it was signaled to wake and check its condition.

*/

func main() {
	c := sync.NewCond(&sync.Mutex{})    // Create condition using standard sync.Mutex as a Locker
	queue := make([]interface{}, 0, 10) // Create a slice with the length of zero.  We will add 10 items eventually so capacity is 10.

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // Enter a critical section for the condition so we can modify data
		queue = queue[1:] // Simulate dequeuing an item by reassigning the head of the slice to the second item.
		fmt.Println("Removed from Queue")
		c.L.Unlock() // Exit the conditions critical section
		c.Signal()   // Here we let a goroutine waiting on a condition know that something has occurred
		// Signal is one of two methods that notifying goroutine.  The other is Broadcast.
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // We enter a critical section for the condition by calling Lock.
		for len(queue) == 2 { // check the length of the queue in the loop.  This is important because the signal
			// on the condition doesn't necessarily mean what you've been waiting for has occurred
			// only that something has occurred
			c.Wait() // Will suspend the main goroutine until a signal on the condition has been sent.
			// Few other things happen when we call Wait: upon entering Wait, Unlock is called on the Cond Locker
			// Upon exiting Wait, the Lock is called on the Cond Locker.
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // Dequeue element after one second using a goroutine
		c.L.Unlock()                        // Here we exit the condition's critical section since we've successfully enqueued an item.
	}
}
