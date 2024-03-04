package main

import (
	"fmt"
	"time"
)

/**
To protect form a `select` that may never unblock we can utilise the Go `time` package.  The `time.After` function
takes in a `time.Duration` argument and returns a channel that will send the current time after the duration you
provide it.  This offers a concise way to time out a select statement.
*/
func main() {
	var c <-chan int
	select {
	case <-c: // case statement will never become unblocked because reading from nil channel.
	case <-time.After(1 * time.Second):
		fmt.Println("Timed Out.")
	}
}
