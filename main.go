package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)

	go setList(channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
