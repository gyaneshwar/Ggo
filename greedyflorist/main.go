package main

import (
	"fmt"

	"github.com/gyaneshwar/Ggo/sorting"
)

func main() {
	fmt.Println(getMinimumCost(3, []int32{1, 3, 5, 7, 9}))
}

// Complete the getMinimumCost function below.
func getMinimumCost(k int32, c []int32) int32 {
	sorting.QuickSortDesc(c) // sorting by ascending order
	// max each will get len(c) % k = bunch size
	bunches := (int32(len(c)) % k)

	var price int32

	// calculate cost each bunch
	for i, j := int32(0), int32(len(c)-1); i <= j; i++ { //todo
		// pick largest, pick smallest till bunch size.
		fmt.Printf("[%d,%d] = %v\n", i, j, c)
		price += c[i]
		fmt.Println(price)
		for k := int32(1); k <= bunches-1 && i < j; k, j = k+1, j-1 {
			price += (k + 1) * c[j]
			fmt.Printf("[%d,%d] = %v\n", i, j, c)
			fmt.Println(price)
		}
	}

	return price

}
