package main

import (
	"fmt"
	"time"
)

/**

Every Go program has at least one goroutine: the main goroutine, which is automatically created and started when the
process begins.

A go routine is a function that is running concurrently along side other code.

They're not OS threads and they're not exactly green threads (threads that are manged by a language's runtime).  They
are high level abstraction known as coroutines.

*/

func main() {
	go sayHello()
	go func() {
		fmt.Println("Woot!!!")
	}()
	someFn := func() {
		fmt.Println("Also works")
	}
	go someFn()
	time.Sleep(1 * time.Second)
}

func sayHello() {
	fmt.Println("Hello")
}
