package main

import "fmt"

/**
The first thing we should do to put channels in right context is to assign channel ownership.  Lets define ownership
of a channel as being a goroutine that instantiates, writes and closes a channel.  Its important to clarify which
goroutine owns a channel in order to reason about our programs logically.

Unidirectional declarations are the tools tht allow us to distinguish between goroutines that own channels and those
that only utilize them: channel owners have write-access view into the chan (`chan` or `chan-<`) and channel utilizers
only have a read-only view into the channel (<-chan).

Once we make this distinction between channel owners and non channel owners, we can assign responsibilities to
goroutines that own channels and those that do not.

Channel owners should:
1. Instantiate the channel
2. Perform writes, or pass ownership to another goroutine
3. Close the channel
4. Encapsulate the previous three things and expose them via a reader channel

Consumer of channels should only need to worry about two things:
1. Know when a channel is closed
2. Responsibly handling blocking for any reason
*/

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // Instantiate the channel, since we know we'll produce six results,
		// we create a buffered channel, so that the goroutine can complete as quickly as possible.
		go func() { // start anonymous goroutine that performs the writes on the resultStream.  Goroutine creation is
			// encapsulated within the surrounding function.
			defer close(resultStream) //  Ensure we close the stream once we are done. As the channel owner, this is our responsibility
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // return the channel as a read-only channel
	}

	resultStream := chanOwner()
	for result := range resultStream { // Read off the channel.  As a consumer, we are only concerned with blocking and closed channels.
		fmt.Printf("Recieved: %d\n", result)
	}
	fmt.Println("Done Receiving!")
}
