package day3

import (
	"fmt"
	"regexp"
)

func part1Sol() int {
	inputArray := ReadInput()

	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}

	for _, line := range inputArray {
		r, _ := regexp.Compile(`(\d{1,},\d{1,})|(\d{1,}x\d{1,})`)
		k := r.FindAllString(line, -1)
		fmt.Println(k)
	}
	return 0
}
