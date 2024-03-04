package main

import (
	"fmt"
	"sync"
)

/**
As the name implies Once is a type that utilizes some `sync` primitives internally to ensure that only one call to `Do`
ever calls the function passed in - even in different goroutines.

Once only counts the number of times `Do` is called, not how many times a unique functions passed into Do are called.

```
var count int
increment := func() { count++ }
decrement := func() { count-- }

var ounce sync.Once
once.Do(increment)
once.Do(decrement)

fmt.Printf("Count: %d\n", count)  // the answer is 1
```

In this way , copies of `Once` are tightly coupled to the functions they are intended to be called with.

*/
func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count) // What is printed here???
}
