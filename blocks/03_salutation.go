package main

import (
	"fmt"
	"sync"
)

/**
Go routine is running a closure that has closed over the iteration variable salutation.  As our loop iterates, salutation
is being assigned to the next string value in the slice literal.

Because the goroutine being scheduled may run at any point in time in the future, it is undetermined what values will
be printed from within the go-routine.  There is a high probability the loop will exit before goroutines are begun.
That means the `salutation` variable falls out of scope. So what happens?

The Go runtime is observant enough to know that a reference to the `salutation` variable is still being held, and therefore
will transfer the memory to the heap so that the goroutines can continue to access it.

Usually on my machine the loop exits before any of the goroutines run, so `salutation` is transferred to the heap
holding a reference to the last value in the string slice.  And so usually see "good day" printed 3 times.

The proper way to write this is to pass salutation into the closure, so by the time the goroutine is running it will be
operating on the data from its iteration.

*/

// what is the expected output
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()

}
