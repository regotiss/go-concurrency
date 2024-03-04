package main

import (
	"context"
	"fmt"
)

/**
In previous example the stages where taking a slice of data and returning a slice of data, i.e batch processing.  They
operate on chunks of data all at once instead of one discrete value at a time.

There is another type of pipeline stage that performs stream processing.  This means that the stage receives and emits
one element at a time.

Channels are uniquely suited to constructing pipelines in Go because the fulfill all of our basic requirements.  They
can receive and emit values, they can safely be used concurrently, they can be ranges over and they are reified by the
language.
*/

// Converts a discrete set of values into a stream of data on a channel.
// At the beginning of a pipeline, you'll have some batch of data that you need to convert to a channel
func generator(ctx context.Context, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-ctx.Done():
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(ctx context.Context, intStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-ctx.Done():
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func add(ctx context.Context, intStream <-chan int, additive int) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-ctx.Done():
				return
			case addedStream <- i + additive:
			}
		}
	}()
	return addedStream
}

func main() {

	intStream := generator(context.Background(), 1, 2, 3, 4)
	pipeline := add(context.Background(), multiply(context.Background(), intStream, 2), 1)

	for v := range pipeline {
		fmt.Println("***", v)
	}

}
