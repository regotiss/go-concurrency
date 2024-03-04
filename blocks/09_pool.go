package main

import (
	"fmt"
	"sync"
)

/**
`Pool` is concurrent-safe implementation of the object pool pattern.

At a high level, the pool pattern is a way to create and make available a fixed number, or pool, of things to use.
Its commonly used to constrain the creation of things that are expensive (eg. database connections) so that only a fixed
number of them are created.

So when working with a Pool, just remember the following points:

- When instantiating `sync.Pool`, give it a `New` member variable that is thread-safe when called.
- When you receive an instance from `Get`, make no assumptions regarding the state of the object you receive back.
- Make sure to call `Put` when you're finished with the object you pulled out of the pool.  Otherwise, the `Pool` is
useless. Usually this is done with `defer`.
- Objects in the pool must be roughly uniform in makeup.

*/

func main() {
	mypool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	mypool.Get()             // These calls invoke the New function defined on the pool since instances haven't yet been instantiated.
	instance := mypool.Get() // Same as above, New function is invoked.
	mypool.Put(instance)     // Here we put the instance back into the pool.  This increases the available number instance to one.
	mypool.Get()             // This call will re-use the instance previously allocated and put it back in the pool.

}
