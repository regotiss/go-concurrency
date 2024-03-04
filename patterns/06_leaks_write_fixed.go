package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
To fix the leak we need to tell producer it can stop.
*/

func main() {
	newRandStream := func(ctx context.Context) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-ctx.Done():
					return
				}
			}
		}()
		return randStream
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	randStream := newRandStream(ctx)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	time.Sleep(2 * time.Second)

}

/******************* Convention Time

Now that we know how to ensure goroutines don’t leak, we can stipulate a convention: If a goroutine is responsible for
creating a goroutine, it is also responsible for ensuring it can stop the goroutine.

********************/
