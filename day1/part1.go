package day1

func part1Sol() int {
	inputArray := ReadInput()

	finalNumber := 0

	for _, num := range inputArray {
		finalNumber += num
	}
	return finalNumber
}
