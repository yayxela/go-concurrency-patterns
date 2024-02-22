package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func main() {
	c := fanIn(generator("First"), generator("Second"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("bye...")
}
