package main

import "fmt"

func duplicate() {
	numbers := [...]int{1, 2, 2, 2, 5, 6, 7, 8, 9, 45}

	for _, v := range numbers {
		if has(numbers, v) >= 2 {
			fmt.Println("Has Duplicates")
			break
		}
	}

}
func has(array [10]int, number int) int {
	i := 0
	for _, v := range array {
		if number == v {
			i++
		}
	}
	return i
}

func main() {

	duplicate()
}
