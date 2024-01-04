package main

func main() {

}

func quickSort(numbers []int) []int {

	if len(numbers) < 2 {
		return numbers
	}

	return quickSort(numbers)
}
