package main

import "fmt"

func main() {
	// result := maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2)
	result := maxJumps([]int{3, 3, 3, 3, 3}, 3)
	fmt.Println(result)
}

type item struct {
	value    int
	index    int
	indices  map[int]bool
	maxjumps int
	assigned bool
}

func maxJumps(arr []int, d int) int {
	sortarr := make([]item, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		sortarr = append(sortarr, item{value: arr[i], index: i, indices: map[int]bool{}})
	}
	QS(sortarr)

	var mj int

	//catching all the possible index jumps
	for i := 0; i < len(arr)-1; i++ {

		j := i
		for {
			j++
			if j >= len(arr) {
				break
			}

			a, b := sortarr[i].index, sortarr[j].index

			left, right := a, b
			if a > b {
				left, right = b, a
			}

			if right-left > d {
				continue
			}

			if !canIJump(arr[left : right+1]) {
				continue
			}

			sortarr[i].indices[b] = true

		}
	}

	//create a map
	mapy := make(map[int]item, len(arr))
	for _, v := range sortarr {
		mapy[v.index] = v
	}

	//find the max
	for i := 0; i < len(sortarr); i++ {
		le := getMaxPathOfAnItem(sortarr[i].index, mapy)
		if le > mj {
			mj = le
		}
	}

	return mj

}

func getMaxPathOfAnItem(index int, items map[int]item) int {
	next := items[index]
	max := 1
	if next.assigned {
		return next.maxjumps
	}

	for nextIndex, _ := range next.indices {
		length := getMaxPathOfAnItem(nextIndex, items) + 1
		if length > max {
			max = length
		}
	}
	next.assigned = true
	next.maxjumps = max
	items[index] = next
	return max
}

func canIJump(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	var min, max int

	if arr[0] == arr[len(arr)-1] {
		return false
	}

	if arr[0] < arr[len(arr)-1] {
		min = arr[0]
		max = arr[len(arr)-1]
	} else {
		max = arr[0]
		min = arr[len(arr)-1]
	}

	//min(i,j) < k < max(i,j)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > min || arr[i] > max {
			return false
		}
	}
	return true
}

func QS(vals []item) {
	if len(vals) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(vals)-1

	vals[pivot], vals[right] = vals[right], vals[pivot]

	for index, v := range vals {
		if v.value > vals[right].value {
			vals[left], vals[index] = vals[index], vals[left]
			left++
		}
	}

	vals[right], vals[left] = vals[left], vals[right]

	QS(vals[:left])
	QS(vals[left+1:])
}
