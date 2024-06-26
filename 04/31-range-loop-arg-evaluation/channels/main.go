package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		close(ch2)
	}()

	ch := ch1
	// When using range with channels, it only provides the value, not an index. This is different from how range works with slices or arrays, where it provides both an index and a value.

	for v := range ch {
		fmt.Println(v)
		ch = ch2
	}
}
