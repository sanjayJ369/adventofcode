package utils

import (
	"strconv"
	"strings"
)

func StringSliceToInts(lines []string) []int {
	res := []int{}
	for _, line := range lines {
		res = append(res, StringToInt(line))
	}
	return res
}

func CombineIntSliceIntoString(nums []int) string {
	var res strings.Builder
	for _, n := range nums {
		res.WriteString(strconv.Itoa(n))
	}
	return res.String()
}

// returns minimumNumber, it's index
func FindMinAndIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, -1
	}

	minNum := 9
	minIdx := -1

	for i, n := range nums {
		if n < minNum {
			minNum = n
			minIdx = i
		}
	}

	return minNum, minIdx
}

func RemoveKDigitsFromSlice(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		nums = RemoveDigitFromSliceToGetLargetNumber(nums)
	}
	return nums
}

// RemoveDigitFromSliceToGetLargetNumber
// remove a digit from a number which is split into []int
// to get the largest number after removing a digit
func RemoveDigitFromSliceToGetLargetNumber(nums []int) []int {
	n := len(nums)
	for i, n := range nums[:n-1] {
		if n < nums[i+1] {
			return append(nums[:i], nums[i+1:]...)
		}
	}

	return nums[:n-1]
}

// returns maximumNumber, it's index
func FindMaxAndIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, -1
	}

	maxNum := nums[0]
	maxIdx := 0

	for i, n := range nums {
		if n > maxNum {
			maxNum = n
			maxIdx = i
		}
	}

	return maxNum, maxIdx
}
