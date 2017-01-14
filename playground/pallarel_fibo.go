package main

import "fmt"

func fibonacci(n uint32, c chan uint32) {
	defer close(c)

	if n == 1 || n == 0 {
		c <- 1
		return
	} else if n < 0 {
		c <- 0
		return
	}
	c_1 := make(chan uint32, 1)
	c_2 := make(chan uint32, 1)
	go fibonacci(n-1, c_1)
	go fibonacci(n-2, c_2)
	c <- <-c_1 + <-c_2
}

func main() {
	var n uint32 = 30
	c := make(chan uint32, 1)
	go fibonacci(n, c)
	fibo := <-c
	fmt.Printf("n: %v, fibonacci: %v\n", n, fibo)
}
