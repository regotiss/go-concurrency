package main

import (
	"context"
	"fmt"
	"math/rand"
)

/**
Some handy generators
*/

func repeat(ctx context.Context, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func take(ctx context.Context, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-ctx.Done():
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toString(ctx context.Context, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-ctx.Done():
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func main() {
	ctx := context.Background()
	for num := range take(ctx, repeat(ctx, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()

	randFn := func() interface{} { return rand.Int() }

	for num := range take(ctx, repeatFn(ctx, randFn), 10) {
		fmt.Println(num)
	}

	var message string
	for token := range toString(ctx, take(ctx, repeat(ctx, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...\n", message)

}
