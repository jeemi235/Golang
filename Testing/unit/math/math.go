package math

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

// go test -cover

// go test -coverprofile coverage.out

// go tool cover -html coverage.out

// go test -bench=Add