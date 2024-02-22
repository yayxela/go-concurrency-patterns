package task

import (
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Response *http.Response
	Err      error
	Num      int
}

func (r *Response) String() string {
	return fmt.Sprintf("{ \"response\": %+v, \"err\": %s, \"num\": %d}", r.Response, r.Err, r.Num)
}

func MakeRequest(url string, num int) <-chan *Response {
	c := make(chan *Response, num)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res, err := http.Get(url)
			c <- &Response{
				Response: res,
				Err:      err,
				Num:      i,
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}
