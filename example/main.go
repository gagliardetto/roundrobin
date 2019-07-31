package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gagliardetto/roundrobin"
)

func main() {
	rrURLs, err := roundrobin.NewURLs([]*url.URL{
		{Host: "192.168.33.10"},
		{Host: "192.168.33.11"},
		{Host: "192.168.33.12"},
		{Host: "192.168.33.13"},
	})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		rrURLs.Next()
		//fmt.Println(rrbn.Next())
	}
	fmt.Println("took:", time.Now().Sub(start))

	///
	rrStrings, err := roundrobin.NewStrings([]string{
		"192.168.33.10",
		"192.168.33.11",
		"192.168.33.12",
		"192.168.33.13",
	})
	if err != nil {
		panic(err)
	}

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		rrStrings.Next()
		//fmt.Println(rrbn.Next())
	}
	fmt.Println("took:", time.Now().Sub(start))
}
