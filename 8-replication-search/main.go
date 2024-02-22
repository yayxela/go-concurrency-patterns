package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var (
	Web          = fakeSearch("web")
	WebReplica   = fakeSearch("web-replica")
	Image        = fakeSearch("image")
	ImageReplica = fakeSearch("image-replica")
	Video        = fakeSearch("video")
	VideoReplica = fakeSearch("video-replica")
)

func Browser(query string) (results []Result) {
	c := make(chan Result)
	//go func() { c <- Web(query) }()
	//go func() { c <- Image(query) }()
	//go func() { c <- Video(query) }()
	go func() { c <- First(query, Web, WebReplica) }()
	go func() { c <- First(query, Image, ImageReplica) }()
	go func() { c <- First(query, Video, VideoReplica) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func main() {
	start := time.Now()
	//fmt.Printf("%v\n", Browser("golang"))
	fmt.Printf("%v\n", Browser("golang"))
	fmt.Println(time.Since(start))
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//var result = First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
