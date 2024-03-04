package main

import (
	"context"
	"fmt"
	"net/http"
)

type Result struct { // we create a type that could hold either a Error or a Response
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(ctx context.Context, urls ...string) <-chan Result {
		responses := make(chan Result)
		go func() {
			defer close(responses)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // no longer swallow error but return a result
				select {
				case <-ctx.Done():
					return
				case responses <- result:
				}
			}
		}()
		return responses
	}

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(context.Background(), urls...) {
		if result.Error != nil { // now we are dealing with the error
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Responses: %v\n", result.Response.Status)
	}
}
