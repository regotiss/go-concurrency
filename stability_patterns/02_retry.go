package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type effectorFn func(context.Context) (string, error)

func retry(effector effectorFn, retries int, delay time.Duration) effectorFn {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}
			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}

func main() {
	var count int

	emulateTransientError := func(ctx context.Context) (string, error) {
		count++

		if count <= 3 {
			return "intentional fail", errors.New("error")
		} else {
			return "success", nil
		}
	}

	r := retry(emulateTransientError, 5, 2*time.Second)

	res, err := r(context.Background())

	fmt.Println(res, err)
}
