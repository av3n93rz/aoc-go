package day2

import "fmt"

func part2Sol() int {
	inputArray := ReadInput()

	for i := 0; i < len(inputArray); i++ {
		for j := i + 1; j < len(inputArray); j++ {
			var charsI []rune
			var charsJ []rune

			for _, value := range inputArray[i] {
				charsI = append(charsI, value)
			}
			for _, value := range inputArray[j] {
				charsJ = append(charsI, value)
			}
			diff := 0
			for index, value := range charsI {
				if value != charsJ[index] {
					diff += 1
				}
			}

			if diff == 1 {
				fmt.Println(inputArray[i])
				fmt.Println(inputArray[j])
			}
		}
	}
	return 0
}
