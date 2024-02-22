package main

import (
	"fmt"
	"time"
)

func generator(done <-chan struct{}, val ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, i := range val {
			select {
			case <-done:
				return
			case c <- i:
			}
		}
	}()
	return c
}

func add(done <-chan struct{}, stream <-chan int, add int) chan int {
	addC := make(chan int)
	go func() {
		defer close(addC)
		for i := range stream {
			select {
			case <-done:
				return
			case addC <- add + i:
			}
		}
	}()
	return addC
}

func multiply(done <-chan struct{}, stream <-chan int, mul int) chan int {
	mulC := make(chan int)
	go func() {
		defer close(mulC)
		for i := range stream {
			select {
			case <-done:
				return
			case mulC <- mul * i:
			}
		}
	}()
	return mulC
}

func main() {
	done := make(chan struct{})
	stream := generator(done, 1, 2, 3, 4, 5)
	pipeline := add(done, multiply(done, stream, 2), 5)
	for {
		select {
		case val := <-pipeline:
			fmt.Println(val)
		case <-time.After(time.Second):
			close(done)
			return

		}
	}
}
