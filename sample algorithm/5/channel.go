package main

import (
	"fmt"
	"math"
)

func checkPrimeNumber(num int, prime chan int, nonprime chan int) {

	if num == 1 {
		// fmt.Printf("%d is Not Prime\n", num)
		nonprime <- num
		return
	}

	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			//fmt.Printf("%d is Not Prime\n", num)
			nonprime <- num
			return
		}
	}
	//fmt.Printf("%d is Prime\n", num)
	prime <- num
	return

}

func checkOddEven(num int, odd chan int, even chan int) {
	if num%2 == 0 {
		//fmt.Printf("%d is odd number\n", num)
		even <- num
	} else {
		//fmt.Printf("%d is even number\n", num)
		odd <- num
	}
}

// Constraints -> length of array < 10 ,maximum number: 1000

func main() {

	prime := make(chan int)
	nonprime := make(chan int)
	odd := make(chan int)
	even := make(chan int)
	fmt.Print("please enter the length of array:")
	var size int
	fmt.Scanln(&size)

	for size > 10 {
		fmt.Print("please enter the length of array less than 10:")
		fmt.Scanln(&size)
	}

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i+1)
		fmt.Scanln(&arr[i])
		for arr[i] > 1000 || arr[i] < 1 {
			fmt.Print("please enter the value less than 1000:")
			fmt.Scanln(&arr[i])
		}
	}

	for i := 0; i < size; i++ {
		num := arr[i]
		go checkPrimeNumber(num, prime, nonprime)
		select {
		case primenumber := <-prime:
			fmt.Printf("%d is prime.\n", primenumber)
		case nonprimenumber := <-nonprime:
			fmt.Printf("%d is not prime.\n", nonprimenumber)
		}
		go checkOddEven(num, odd, even)
		select {
		case oddnumber := <-odd:
			fmt.Printf("%d is odd.\n", oddnumber)
		case evennumber := <-even:
			fmt.Printf("%d is even.\n", evennumber)
		}
	}
}
