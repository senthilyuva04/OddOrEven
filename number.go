package main

import "fmt"

func main() {
	var number = NewNumbers()
	for _, n := range number {
		if n%2 == 0 {
			fmt.Println(n, " is even")
			continue
		}
		fmt.Println(n, " is odd")
	}
}

func NewNumbers() [10]int {
	var numbers [10]int
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}
	return numbers
}
