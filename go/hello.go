package main

import (
	"fmt"
)

func multiplyByTo(in <-chan int, out chan<- int) {
	num := <-in
	result := num * 2
	out <- result
}

func main() {
	out := make(chan int)
	in := make(chan int)

	go multiplyByTo(in, out)
	go multiplyByTo(in, out)

	in <- 1
	in <- 2

	fmt.Println(<-out)
	fmt.Println(<-out)
}