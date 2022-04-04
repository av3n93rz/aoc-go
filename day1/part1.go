package main

func part1Sol() int {
	inputArray := readInput()

	finalNumber := 0

	for _, num := range inputArray {
		finalNumber += num
	}
	return finalNumber
}
