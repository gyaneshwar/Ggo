package main

import (
	"fmt"

	"github.com/gyaneshwar/Ggo/sorting"
)

/*
 * Complete the 'minimumAbsoluteDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func minimumAbsoluteDifference(arr []int32) int32 {
	sorting.QuickSort(arr)
	var min int32 = 0
	for i := 0; i < len(arr)-1; i++ {
		if i == 0 {
			min = arr[i] - arr[i+1]
			if min < 0 {
				min *= -1
			}
			continue
		}

		tm := arr[i] - arr[i+1]
		if tm < 0 {
			tm *= -1
		}

		if tm < min {
			min = tm
		}
	}

	return min
}

func main() {
	fmt.Println(minimumAbsoluteDifference([]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}))
}
