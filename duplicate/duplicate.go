package main

import (
	"fmt"
	"sort"
)

func duplicate() {
	numbers := []int{9, 5, 13, 2, 4, 6, 7, 8, 1, 3}
	sort.Ints(numbers)
	for i := 0; i < len(numbers); i++ {
		if i != len(numbers)-1 {
			if numbers[i] == numbers[i+1] {
				fmt.Println("Has Duplicates")
				break
			}
		}

	}
}

func main() {

	duplicate()
}
