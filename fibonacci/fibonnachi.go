package main

import "fmt"

type Fibonacci struct{}

func fibonacci(size int) {
	first, second := 0, 1
	for i := 0; i < size; i++ {
		fmt.Printf("%v ", first)
		first, second = second, first+second
	}
}
func main() {
	var number int
	fmt.Scan(&number)

	fibonacci(number)
}
