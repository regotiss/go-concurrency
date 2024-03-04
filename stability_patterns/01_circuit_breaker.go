package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

/**
Circuit Breaker automatically degrades service functions in response to a likely fault, preventing larger or cascading
failures by eliminating recurring errors and providing reasonable error responses.
*/

/**
Define the signature of the function interacting with you upstream service.
*/
type circuitFn func(context.Context) (string, error)

func Breaker(circuit circuitFn, failureThreshold uint) circuitFn {
	var consecutiveFailures = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() //establish a read lock
		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock()                   // Release the read lock
		response, err := circuit(ctx) // Issue request

		m.Lock() // lock around the shared resource
		defer m.Unlock()

		lastAttempt = time.Now() // Record time of the attempt
		if err != nil {          // Circuit returned an error, so we count the failures and return
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0 // reset failures
		return response, nil
	}
}

func main() {
	attempts := 0
	doSomething := func(ctx context.Context) (string, error) {
		if attempts < 3 {
			attempts++
			return "", errors.New("500 Service unavailable")
		}
		return "It Works!!", nil
	}

	fn := Breaker(doSomething, 2)

	for {
		v, err := fn(context.Background())
		if err == nil {
			fmt.Println(v)
			break
		} else {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}

}
