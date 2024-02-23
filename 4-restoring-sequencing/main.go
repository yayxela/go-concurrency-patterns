package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func generator(msg string) <-chan Message {
	c := make(chan Message)
	go func() {
		wait := make(chan bool)
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: wait,
			}
			<-wait
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan Message) <-chan Message {
	c := make(chan Message)
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
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("bye...")
}
