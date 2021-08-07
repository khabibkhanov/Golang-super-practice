package main

import "fmt"

func fizzBuzz(lastNum int) {

	for i := 1; i <= lastNum; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Print(" FizzBuzz ")
		} else if i%5 == 0 {
			fmt.Print(" Buzz ")
		} else if i%3 == 0 {
			fmt.Print(" Fizz ")
		} else {
			fmt.Print(i)
		}
	}
}

func main() {
	var number int
	fmt.Scan(&number)
	fizzBuzz(number)
}
