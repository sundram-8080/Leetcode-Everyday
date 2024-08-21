package problems

import (
	"fmt"
)

func ContainsDuplicate(testcase []int) bool {

	maps := make(map[int]bool)

	for index, value := range testcase {

		if maps[value] {

			fmt.Println("Found doplicate at ", index)
			return true
		}

		maps[value] = false

	}

	return false
}
