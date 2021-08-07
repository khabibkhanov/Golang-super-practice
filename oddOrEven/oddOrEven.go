package main

import "fmt"

func oddAndEven() {
	var oddSum, evenSum int
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			oddSum += i
		} else {
			evenSum += i
		}
	}
	fmt.Printf("\n Even sum: %v \n Odd sum: %v \n", evenSum, oddSum)
}

func main() {

	oddAndEven()
}
