package day2

func part2Sol() string {
	inputArray := ReadInput()
	res := ""

	for i := 0; i < len(inputArray); i++ {
		for j := i + 1; j < len(inputArray); j++ {
			var charsI []string
			var charsJ []string

			for _, value := range inputArray[i] {
				charsI = append(charsI, string(value))
			}
			for _, value := range inputArray[j] {
				charsJ = append(charsJ, string(value))
			}
			diff := 0

			for index, value := range charsI {
				if value != charsJ[index] {
					diff += 1
				}
			}

			if diff == 1 {
				similar := inputArray[j]
				for index, value := range inputArray[i] {
					current := string(value)
					if string(similar[index]) == current {
						res += current
					}
				}
			}
		}
	}
	return res
}
