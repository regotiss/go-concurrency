package main

import "fmt"

/**
A pipeline is just another tool you can use to form an abstraction in your system.  In particular, it is very powerful
tool to use when your program needs to process streams, or batches of data.

A pipeline is nothing more than a series of things that take data in, perform an operation on it, and pass the data
back out.  We call each of these operations a stage of the pipeline.

By using a pipeline, you separate the concerns of each stage, which provides numerous benefits.  You can modify stages
independent of one another, you can mix and match how stages are combined independent of modifying the stages.

Properties of a pipeline:
- A stage consumes and returns the same type
- A stage must be reified by the language so that it may be passed around.  Function in Go are reified and fit this
  purpose nicely.

People that are familiar with functional programming are probably thinking higher order functions or monads.  Pipelines can be
considers a subset of monads.
*/

func main() {
	// function can be considered a pipeline stage
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	// another pipeline stage
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	// now we can combine them.
	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}

	// can write the same by doing this
	for _, v := range ints {
		fmt.Println(v*2 + 1)
	}
}
