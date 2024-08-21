package problems

import (
	"sort"
)

func TopKfrequentElements(arr []int, k int) []int {
	maps := make(map[int]int)

	for _, val := range arr {
		maps[val]++
	}

	var countArr [][]int

	for key, value := range maps {
		countArr = append(countArr, []int{key, value})
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i][1] > countArr[j][1]
	})

	var result []int

	for i := 0; i < k; i++ {
		result = append(result, countArr[i][0])
	}

	return result
}
