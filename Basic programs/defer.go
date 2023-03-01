package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		defer fmt.Print(i, " ")
	}
}

// func main()  {

//     for i:= 1; i <= 10; i++ {
//         defer fmt.Print(i, " ")
//     }
// }
