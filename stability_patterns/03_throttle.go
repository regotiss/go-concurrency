package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type throttleEffector func(context.Context) (string, error)

func throttle(e throttleEffector, max uint, refill uint, d time.Duration) throttleEffector {
	var tokens = max
	var once sync.Once

	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return "", fmt.Errorf("too many calls")
		}

		tokens--

		return e(ctx)
	}
}

func main() {

	fn := func(ctx context.Context) (string, error) {
		fmt.Println("Called Function!!")
		return "string", nil
	}

	throttledFn := throttle(fn, 5, 5, time.Second*5)

	attempt := 0
	for {
		_, err := throttledFn(context.Background())
		if err != nil {
			fmt.Println(err)
			attempt++
			fmt.Println("Waiting 6 seconds!!")
			time.Sleep(6 * time.Second)
			if attempt > 2 {
				break
			}
		}
	}

}
