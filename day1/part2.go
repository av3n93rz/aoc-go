package day1

func part2Sol() int {

	m := make(map[int]bool)

	inputArray := ReadInput()

	finalNumber := 0
	result := false

	for !result {
		for _, num := range inputArray {
			finalNumber += num

			if m[finalNumber] {
				result = true
				break
			}
			m[finalNumber] = true
		}
	}
	return finalNumber
}
