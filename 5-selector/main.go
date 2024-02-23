package main

import (
	"fmt"
)

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c1 := generator("First")
	c2 := generator("Second")
	select {
	case v1 := <-c1:
		fmt.Println(v1)
	case v2 := <-c2:
		fmt.Println(v2)
		//default:
		//	fmt.Println("no one is ready")
	}
}
