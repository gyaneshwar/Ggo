package main

import (
	"fmt"
)

//https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/submissions/
func smallerNumbersThanCurrent(nums []int) []int {
	//make a copy of the original array.
	originalNums := make([]int, len(nums))
	copy(originalNums, nums)

	//sort the array
	QS(nums)

	for k, v := range originalNums {
		position := 0
		for i := 0; i < len(nums); i++ {
			//once the item is matched, find the last position of the matched item (eg: 1,2,2 . position of 2 is 3.)
			if nums[i] == v {
				position = i
				for j := i + 1; j < len(nums); j++ {
					if nums[j] != v {
						break
					}
					position++
				}
				break
			}
		}

		// find the length from the position of the item till the end.
		originalNums[k] = len(nums) - position - 1
	}
	return originalNums
}

func QS(a []int) {
	if len(a) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(a)-1

	a[pivot], a[right] = a[right], a[pivot]

	for i, v := range a {
		if v > a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[right], a[left] = a[left], a[right]

	QS(a[:left])
	QS(a[left+1:])
}

func main() {
	result := []int{4, 0, 1, 1, 3}
	answer := smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})
	for i, a := range answer {
		if result[i] == a {
			continue
		}
		fmt.Printf("Error")
	}
	fmt.Printf("%v", answer)
}
