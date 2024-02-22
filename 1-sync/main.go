package main

func main() {
	// каналы
	c := make(chan int)
	c <- 1
	val := <-c
	_ = val
}
