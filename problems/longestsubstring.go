package problems

import (
	"fmt"
)

func Longestsubstring(input string) {
	fmt.Println("input is : ", input)

	seen := make(map[rune]int)

	l := 0

	length := 0

	for r, val := range input {

		if _, ok := seen[val]; ok && r > l {

			fmt.Println("found", val)

			l = seen[val] + 1

			fmt.Println(seen)
		} else {
			currentLength := r - l + 1

			if currentLength > length {
				length = currentLength
			}
		}

		seen[val] = r

	}

}
