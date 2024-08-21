package main

import (
	"fmt"

	"netnat.xyz/leetcode/problems"
)

func main() {

	GroupAnagram()

}

func GroupAnagram() {

	testcases := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	result := problems.GroupAnagram(testcases)

	fmt.Println(result)
}

func twosum() {
	nums := []int{5, 3, 2, 7, 11, 15}
	target := 9

	result := problems.TwoSum(nums, target)

	fmt.Println(result)
}

func longestsubstring() {

	input1 := "abcabcdd"
	// input2 := "pwwkew"
	// input3 := "abcacbdd"

	problems.Longestsubstring(input1)
}

func containsDuplicate() {

	arr := []int{1, 2, 3, 2, 4, 5, 1}

	result := problems.ContainsDuplicate(arr)

	fmt.Println(result)
}

func validAnagram() {
	s := "anagram"
	t := "nagaram"

	result := problems.ValidAnangram(s, t)

	fmt.Println(result)
}
