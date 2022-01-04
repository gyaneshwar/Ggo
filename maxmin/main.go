package main

import (
	"fmt"

	"github.com/gyaneshwar/Ggo/sorting"
)

func main() {
	fmt.Println(maxMin(3, []int32{10, 100, 300, 200, 1000, 20, 30}))
}

func maxMin(k int32, arr []int32) int32 {
	// Write your code here
	sorting.QuickSortDesc(arr)
	var minDiff int32 = arr[0]
	for i := int32(0); i <= int32(len(arr))-k; i++ {
		md := arr[i] - arr[i+k-1]
		fmt.Printf("[%d,%d] %v = %d\n", i, i+k-1, arr, md)

		if md < minDiff {
			minDiff = md
		}
	}

	return minDiff
}
