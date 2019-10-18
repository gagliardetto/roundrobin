package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gagliardetto/roundrobin"
)

func main() {
	////
	rrURLs, err := roundrobin.New([]interface{}{
		&url.URL{Host: "192.168.33.10"},
		&url.URL{Host: "192.168.33.11"},
		&url.URL{Host: "192.168.33.12"},
		&url.URL{Host: "192.168.33.13"},
	})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for i := 0; i < 10; i++ {
		u := rrURLs.Next()
		uu := u.(*url.URL)
		fmt.Println(uu.Host)
		uu.Host = "modified"
		//fmt.Println(rrbn.Next())
	}
	fmt.Println("took:", time.Now().Sub(start))
	return

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
