package main

import "fmt"

/**
Best used to communicate information between goroutine.

Like a river, a channel serves as a conduit for a stream of information; values may be passed along the channel, and
then read out downstream.  When using channels, you'll pass a value into a `chan` variable, and then somewhere else in
your program read it off the channel.  The disparate parts of you program don't require knowledge of each other, only a
reference to the same place in memory where the channel resides.  This can be done by passing references of channels
around your program.

Channels are typed.  Can create a channel `chan interface{}` which means that we can place any kind of data onto it, but
you can also give it a stricter type to constrain it.

To declare a unidirectional channel, you'll simply include the `<-` operator.  To both declare and instantiate a channel
that can only read, place the `<-` operator on the lefthand side.

```
var dataStream <-chan interface{}
datastream := make(<-chan interface{})
```

To declare and create a channel that can only send, you place `<-` operator on the righthand side

```
var dataStream chan<- interface{}
dataStream := make(chan<- interface{})
```

You don't often see unidirectional channels instantiated, but you'll often see them used as function parameters and
return types, which is very useful.  This is possible because Go will implicity convert bidirectional channels to
unidirectional channels when needed.

```
var receiveChan <-chan interface{}
var sendChan chan<- interface{}
dataStream := make(chan interface{})

receiveChan = dataStream
sendChan = dataStream
```

*/

func main() {
	var receiveChan <-chan string
	var sendChan chan<- string
	stringStream := make(chan string)
	receiveChan = stringStream
	sendChan = stringStream
	go func() {
		sendChan <- "Hello channels!" // Put something onto the channel
	}()
	fmt.Println(<-receiveChan) // Pull something off the stream

}
