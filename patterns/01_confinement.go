package main

import (
	"bytes"
	"fmt"
	"sync"
)

/**
  When working with concurrent code, we've gone over two ways to ensure safe operation
    - Synchronization primitives for sharing memory (eg. sync.Mutex)
    - Synchronization via communicating (eg. Channels)

  However there a couple of other options that are implicitly safe:
    - Immutable data
    - Data protected by confinement

  In some sense, immutable data is ideal because is implicitly concurrent-safe.  Each concurrent process may operate
  on the same data, but it may not modify it.  If it wants to create new data, it must make a copy with the desired
  modifications.  In Go, you can achieve this by writing code that utilizes copies of values instead of pointers to
  values in memory.

  Confinement can also allow for a lighter cognitive load on the developer and smaller critical sections.  Confinement
  is the simple yet powerful idea of ensuring information is only ever available from one concurrent process.  When this
  is achieved, a concurrent program is implicitly safe and no synchronization is needed.  There are two kinds of
  confinement ad hoc and lexical.

  Ad hoc confinement is when you confinement through a convention - whether it be set by the language community, the
  group you work with, or the codebase you work with.  Sticking to convention is difficult to achieve on projects of any
  size unless you have tools to perform static analysis on your code every time someone commits some code.

*/

func adhoc() {
	data := make([]int, 4) // this is available from both the `loopData` function and the loop over the `handleData` channel;
	// however by convention we're only accessing it from the `loopData` function.

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

/**
	Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives for
    multiple concurrent processes to use.  We've touched upon this previously when discussing channel ownership.
*/

func lexical() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // Here we instantiate the  channel within the lexical scope of `chanOwner` function
		// This limits the scope of the write aspect fo the results channel to the closure defined below it.  It confines the write
		// aspect of this channel to prevent other goroutines from writing to it.
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // Here we receive a read-only copy of an int channel.  By declaring that only
		// usage we require is read access, we confine usage of the channel within the `consume` function to only reads.
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // Here we receive the read aspect of the channel and we're able to pass it into the `consumer`
	// which can do nothing but read from it.  Once again this confines the main goroutine to a read-only view of the channel.
	consumer(results)
}

/**
  Since channels are concurrent-safe, lets look at an example of confinement that uses a data structure which is not
  currently safe, an instance of bytes.Buffer.
*/

func databytes() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // pass a slice containing the first 3 bytes
	go printData(&wg, data[3:]) // pass a slice containing the last 3 bytes

	wg.Wait()
}

/**

In the above example, you can see that because printData doesn’t close around the data slice, it cannot access it, and needs
to take in a slice of byte to operate on. We pass in different subsets of the slice, thus constraining the goroutines
we start to only the part of the slice we’re passing in. Because of the lexical scope, we’ve made it impossible to do
the wrong thing, and so we don’t need to synchronize memory access or share data through communication.

So what’s the point? Why pursue confinement if we have synchronization available to us? The answer is improved
performance and reduced cognitive load on developers. Synchronization comes with a cost, and if you can avoid it
you won’t have any critical sections, and therefore you won’t have to pay the cost of synchronizing them. You also
sidestep an entire class of issues possible with synchronization; developers simply don’t have to worry about these
issues. Concurrent code that utilizes lexical confinement also has the benefit of usually being simpler to understand
than concurrent code without lexically confined variables. This is because within the context of your lexical scope you
can write synchronous code.

Having said that, it can be difficult to establish confinement, and so sometimes we have to fall back to our wonderful
Go concurrency primitives.

*/

func main() {
	adhoc()
	lexical()
	databytes()
}
