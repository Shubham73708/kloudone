// 1)Creating Concurrent Workflows in Go with Goroutines and Channels

package main

import (
	"fmt"
)

func evenNum(even chan int) {
	i := 2
	for i < 10 {
		even <- i
		i = i + 2
	}
	close(even)
}

func oddNum(odd chan int) {
	i := 1
	for i < 10 {
		odd <- i
		i = i + 2
	}
	close(odd)
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	go evenNum(even)
	go oddNum(odd)
	for {
		even, ok1 := <-even
		odd, ok2 := <-odd
		if ok1 == false && ok2 == false {
			break
		}
		fmt.Println("Received ", even, ok1, odd, ok2)
	}
}
