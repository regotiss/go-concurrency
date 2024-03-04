package main

import (
	"bytes"
	"fmt"
	"os"
)

/**
We can also create buffered channels, which are channels that are given a capacity when they're instantiated.
This means that even if no reads are performed on channels, a goroutine call still perform `n` writes, where `n` is the
capacity of the buffered channel.

We use the `make` function to determine the capacity of the channel, which is interesting because it means that the
goroutine that instantiates a channel controls whether its buffered.  This suggests that the creation of the channel
should probably be closely coupled to goroutines that will be performing writes on it so that we can reason about its
behaviour and performance easily.

An unbuffered channel is simply a buffered channel with a capacity of 0.
*/

func main() {
	var stdoutBuff bytes.Buffer         // Create an in-memory buffer
	defer stdoutBuff.WriteTo(os.Stdout) // Ensure buffer written to stdout before process exits.

	intStream := make(chan int, 4) // Create a buffered channel.
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
