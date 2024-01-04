package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(count(numbers))
}

func count(numbers []int) int {
	if len(numbers) <= 1 {
		return 1
	}

	return 1 + count(numbers[1:])
}
