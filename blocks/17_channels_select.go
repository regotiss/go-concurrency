package main

import (
	"fmt"
	"time"
)

/**
The `select` statement is the glue that binds channels together; it's how we're able to compose channels together
in a program to form a larger abstraction.

`select` statements can help safely bring channels together with concepts like cancellations, timeouts, waiting and
default values.

It looks a bit like a switch block.  Just like a switch block, a `select` block encompasses a series of `case` statements
that guard a series of statements; however, that where the similarities end.  Unlike `switch` blocks, `case` statements
in a `select` block aren't tested sequentially, and execution won't automatically fall through if none of the criteria are
met.

Instead, all channel reads and writes are considered simultaneously to see if any of them are ready: populated or closed
channels in case of reads, and channels that are not at capacity in the case of writes.  If none of the channels are ready,
the entire `select` statement blocks.  Then when one of the channels are ready, that operation will proceed, and its
corresponding statements will execute.
*/

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // Close after waiting for five seconds.
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: //Attempt to read on the channel
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
