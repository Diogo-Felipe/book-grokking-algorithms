package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers))
}

func sum(numbers []int) int {
	if len(numbers) <= 1 {
		return numbers[0]
	}

	return numbers[0] + sum(numbers[1:])

}
