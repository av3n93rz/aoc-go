package day2

func part1Sol() int {
	inputArray := ReadInput()

	twiceCount := 0
	thriceCount := 0

	for _, line := range inputArray {

		m := make(map[string]int)
		twice := false
		thrice := false

		for _, rune := range line {
			char := string(rune)
			if m[char] != 0 {
				m[char] += 1
			} else {
				m[char] = 1
			}
		}

		for _, value := range m {
			if value == 2 {
				twice = true
			}
			if value == 3 {
				thrice = true
			}
		}
		if twice {
			twiceCount += 1
		}

		if thrice {
			thriceCount += 1
		}
	}
	return twiceCount * thriceCount
}
