package task

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	num := 3
	res := MakeRequest("https://httpbin.org/get", num)
	list := make([]*Response, 0)
	for r := range res {
		fmt.Println(r)
		list = append(list, r)
	}
	require.Equal(t, len(list), num)
}
