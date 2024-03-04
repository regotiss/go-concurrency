package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

/**

Critical sections are usually reflect a bottle neck in your program, it can be expensive to enter and exit so we try
and minimise the time spent in critical sections.

There may be memory that needs to be shared between multiple concurrent processes, but perhaps not all of these process
will read and write to this memory. If this is the case we can use sync.RWMutex.

*/

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) { // sync.Locker is an interface with two methods, Lock and Unlock,
		// which Mutex and RWMutex satisfy.
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) // make the writer sleep for a second to make the producer less active then the observer
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m))
	}
}
