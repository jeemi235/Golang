package main

import "fmt"

func Compare(a [3]int, b [2]int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func main() {
	a := [3]int{1, 3, 1}
	b := [2]int{1, 3}

	var res int
	res = Compare(a, b)

	fmt.Println(res)
}
