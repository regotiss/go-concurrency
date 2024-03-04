package main

import (
	"context"
	"fmt"
	"time"
)

/**

Something you’ll see over and over again in Go programs is the for-select loop. It’s nothing more than something like this:

```
  for { // Either loop infinitely or range over something
    select {
    //Do some work with channels
   }
  }

```

There are a couple of different scenarios where you'll see this pattern pop up.

  - Sending iteration variables out on  a channel
    Oftentimes you'll want to convert something that can be iterated over into values on a channel.  This is nothing
    fancy
*/

func func1(context context.Context) <-chan string {
	stringStream := make(chan string, 3)

	go func() {
		defer close(stringStream)
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-context.Done():
				fmt.Println("Done()!!")
				return
			case stringStream <- s:
			}
		}
	}()

	return stringStream
}

/**
Looping infinitely waiting to be stopped

 - Its very common to create goroutines that loop infinitely until they're stopped.  There are a couple variations
   of this one.  Which one you choose is purely stylistic preference
*/
func func2(context context.Context) {
	for {
		select {
		case <-context.Done():
			return
		default:
			fmt.Println("doing some work")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	readStream := func1(context.Background())
	for v := range readStream {
		fmt.Println(v)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	func2(ctx)
}
