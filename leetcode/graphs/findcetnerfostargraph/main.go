// 7:01 - 7:13 https://leetcode.com/problems/find-center-of-star-graph/

package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {3, 1}}))
}

func findCenter(edges [][]int) int {
	counter := map[int]int{}
	for _, edge := range edges {
		for _, e := range edge {
			if v, ok := counter[e]; !ok {
				counter[e] = 1
			} else {
				v++
				counter[e] = v
			}
		}
	}

	result := []int{}
	for key, count := range counter {
		if count == len(counter)-1 {
			result = append(result, key)
		}
	}

	if len(result) == 1 {
		return result[0]
	}

	return -1
}
