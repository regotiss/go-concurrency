package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(ctx context.Context, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err) // Swallows the error, is there a better way??
					continue
				}
				select {
				case <-ctx.Done():
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(context.Background(), urls...) {
		fmt.Printf("Responses: %v\n", response.Status)
	}
}
