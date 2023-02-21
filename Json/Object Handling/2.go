package main

import "fmt"

func main() {
	types := []interface{}{"a", 6, 6.0, true}
	for _, v := range types {
		fmt.Printf("%T\n", v)
	}
}
