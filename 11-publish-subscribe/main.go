package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Client interface {
	PublishMessage(msg string)
	ReadMessage() <-chan string
	Close()
}

type client struct {
	c chan string
}

func NewClient() Client {
	return &client{c: make(chan string)}
}

func (c *client) PublishMessage(msg string) {
	c.c <- msg
}

func (c *client) ReadMessage() <-chan string {
	return c.c
}

func (c *client) Close() {
	close(c.c)
}

func main() {
	pubSub := NewClient()
	go func() {
		for msg := range pubSub.ReadMessage() {
			fmt.Println(msg)
		}
	}()

	for i := 0; i < 5; i++ {
		pubSub.PublishMessage(fmt.Sprintf("new message at %s\n", time.Now().Format(time.RFC3339)))
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	pubSub.Close()
}
