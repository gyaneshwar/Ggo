package main

import (
	"fmt"

	"github.com/gyaneshwar/Ggo/sorting"
)

func main() {
	fmt.Println(luckBalance(3, [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}))
}

// luck balance, here the contestent needs to win on her will the easiest games, so that she can maximize her winnings by using luck.
// the problem is solved by choosing the easiest one's first in the must win category (contest1s)
// and then she could maximize her winnings by luck in rest of the tough games.
func luckBalance(k int32, contests [][]int32) int32 {

	// seperate 1's and 0's
	contests1s := []int32{}
	contests0s := []int32{}

	for _, v := range contests {
		if v[1] == 1 {
			contests1s = append(contests1s, v[0]) //[5,2,1,8]
			continue
		}
		contests0s = append(contests0s, v[0]) // [5,10]
	}

	// sort by luck factor desc for contents1s
	sorting.QuickSort(contests1s) // [1,2,5,8]
	fmt.Println(contests1s)

	var sumK, sumAll int32
	var winsWithOutLuck = int32(len(contests1s)) - k
	// add top k values from contents 1s , and add rest of the remaining from contents1s and contest0s
	for i := int32(0); i < int32(len(contests1s)); i++ {
		if winsWithOutLuck > i {
			sumK += contests1s[i] // 1
			fmt.Println(sumK)
			continue
		}

		sumAll += contests1s[i] // 2 + 5 + 8
		fmt.Println(sumAll)
	}

	for _, v := range contests0s {
		sumAll += v // 5 + 10
		fmt.Println(sumAll)
	}

	return sumAll - sumK
}
