package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}


// // Go program to illustrate how
// // to close a channel using for
// // range loop and close function
// package main

// import "fmt"

// // Function
// func myfun(mychnl chan string) {

// 	for v := 0; v < 4; v++ {
// 		mychnl <- "GeeksforGeeks"
// 	}
// 	close(mychnl)
// }

// // Main function
// func main() {

// 	// Creating a channel
// 	c := make(chan string)

// 	// calling Goroutine
// 	go myfun(c)

// 	// When the value of ok is
// 	// set to true means the
// 	// channel is open and it
// 	// can send or receive data
// 	// When the value of ok is set to
// 	// false means the channel is closed
// 	for {
// 		res, ok := <-c
// 		if ok == false {
// 			fmt.Println("Channel Close ", ok)
// 			break
// 		}
// 		fmt.Println("Channel Open ", res, ok)
// 	}
// }
