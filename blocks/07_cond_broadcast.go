package main

import (
	"fmt"
	"sync"
)

/**
Internally, the runtime maintains a FIFO list of goroutines waiting to be signaled; Signal find the goroutine
that's been waiting longest and notifies that, whereas Broadcast sends a signal to all goroutines that are waiting.

Cond is much more performant then using channels.

Like most other things in the sync package, usage of Cond works best when constrained to a tight scope, or exposed
to a broader scope through a type that encapsulates it.
*/

type Button struct { // Define a button with a condition, Clicked.
	Clicked *sync.Cond
}

func main() {

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { // Convenience function that allows us to register functions to handle
		// signals from a condition.  Each handler is run in its own goroutine, and subscribe wont exit until that
		// goroutine is confirmed to be running.

		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // Set up a wait group to ensure program doesn't exit until we write to stdout.
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // simulate the user has clicked the button

	clickRegistered.Wait()

}
