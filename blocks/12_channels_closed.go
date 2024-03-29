package main

import "fmt"

/**
The receiving operator `<-` can also optionally return two values.  The value put on the channel and a boolean value.

What does the boolean signify?  the second return value is a way for a read operation to indicate whether the read off
the channel was a value generated by a write elsewhere in the process, or a default value generated from a closed channel.

It's useful to indicate that no more values will be sent over the channel.  This helps downstream processes know when
to move on, exit, re-open communications on a new or different channel etc.  To close a channel, we use the `close`
keyword.

Even after a channel has been closed you can continue to read from it indefinitely, the second value will indicate that
the value received is the zero value and not a value placed on the stream.
*/

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", salutation, ok)
	close(stringStream)
	salutation, ok = <-stringStream // reading from a closed stream.
	fmt.Printf("(%v): %v\n", salutation, ok)
}
