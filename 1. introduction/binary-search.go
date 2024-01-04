package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	position, found := search(numbers, 3)

	if found {
		fmt.Println("Found at position ", position)
	} else {
		fmt.Println("Element not found")
	}
}

func search(numbers []int, searchNumber int) (int, bool) {
	lower := 0
	high := len(numbers) - 1

	for lower <= high {
		middle := (high - lower) / 2
		middleValue := numbers[middle]

		if searchNumber == middleValue {
			return middle, true
		}

		if searchNumber > middleValue {
			high = middle - 1
		} else {
			lower = middle + 1
		}
	}

	return -1, false
}
