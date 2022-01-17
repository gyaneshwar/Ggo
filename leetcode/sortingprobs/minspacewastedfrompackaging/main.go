// https://leetcode.com/problems/minimum-space-wasted-from-packaging/  8:38 - 9:32
package main

import (
	"fmt"
)

func main() {
	//space := minWastedSpace([]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}, {3, 4}})
	space := minWastedSpace([]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}})
	if space == 9 {
		fmt.Println("awesome")
		return
	}
	fmt.Println(space)
	fmt.Println("failure")
}

type provider struct {
	bxSizes []int //sort the sizes
	space   int
}

func (p *provider) measureSpace(pk []int) {
	//take a box,
	// try to fit a package, if not possible, move to next box, once fixed, get teh difference and add to space
	pkPosition := 0
	bxIndex := 0
	for ; bxIndex < len(p.bxSizes); bxIndex++ {

		for ; pkPosition < len(pk); pkPosition++ {

			if p.bxSizes[bxIndex] >= pk[pkPosition] {
				p.space += p.bxSizes[bxIndex] - pk[pkPosition]
				continue
			}
			break
		}
	}

	if len(pk) != pkPosition {
		p.space = -1
		return
	}

	p.space = p.space % (10e9 + 7)
}

func minWastedSpace(packages []int, boxes [][]int) int {

	//arrange packages in asc.
	QS(packages)

	spaces := []int{}

	for _, v := range boxes {
		//arrange provider boxes in asc
		QS(v)
		p := provider{bxSizes: v}
		p.measureSpace(packages)
		spaces = append(spaces, p.space)
	}

	//find the minimum positive size
	var result int
	var setPositive = false
	for _, space := range spaces {
		if space > 0 && !setPositive {
			setPositive = true
			result = space
			continue
		}
		if space < 0 {
			continue
		}
		if space < result {
			result = space
		}
	}

	// if none of the values are positives, return -1 (boxes not found)
	if !setPositive {
		return -1
	}

	return result
}

func QS(vals []int) {
	if len(vals) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(vals)-1

	vals[pivot], vals[right] = vals[right], vals[pivot]

	for index, v := range vals {
		if v < vals[right] {
			vals[left], vals[index] = vals[index], vals[left]
			left++
		}
	}

	vals[right], vals[left] = vals[left], vals[right]

	QS(vals[:left])
	QS(vals[left+1:])
}
