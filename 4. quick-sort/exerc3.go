package main

import "fmt"

func main() {
	numbers := []int{1, 4, 2, 8, 3, 6, 5, 9, 7}
	fmt.Println(highestValue(0, numbers))
}

func highestValue(previousHighest int, numbers []int) int {

	actualHighest := previousHighest

	if numbers[0] > previousHighest {
		actualHighest = numbers[0]
	}

	if len(numbers) <= 1 {
		return actualHighest
	}

	return highestValue(actualHighest, numbers[1:])
}
