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
			time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := generator("One")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(time.Second):
			fmt.Println("too slow. bye...")
			return
		}
	}
}
