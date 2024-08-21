package problems

import "fmt"

func TwoSum(nums []int, target int) []int {
	checked := make(map[int]int)

	for index, num := range nums {

		if val, ok := checked[target-num]; ok {
			fmt.Println("Found the match at ", val, index)
			return []int{val, index}
		} else {
			fmt.Println("Didn't found the match at ", index)
			checked[num] = index
		}
	}

	return []int{}
}
