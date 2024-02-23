package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		quit <- "bye"
		fmt.Println(<-quit)
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := generator("One", quit)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			quit <- "see ya"
		}
	}

}
