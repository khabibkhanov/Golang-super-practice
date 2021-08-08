package main

import "fmt"

func polindrone(text string) string {
	return text + Reverse(text[:len(text)-1])
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {

	var text string
	fmt.Scan(&text)
	fmt.Println(polindrone(text))
}
