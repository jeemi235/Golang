package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := 1.0
	for n := 0; n < 10; n++ {
		res = res - ((res*res - x) / (2 * res))
	}
	return res
}

func main() {
	z := 169.0
	fmt.Println(Sqrt(z))
	fmt.Println(math.Sqrt(z))
}
