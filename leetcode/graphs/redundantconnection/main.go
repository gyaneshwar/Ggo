// https://leetcode.com/problems/redundant-connection/ - 8:26 , submitted 8:50 ,
// solution and possible approaches. https://leetcode.com/problems/redundant-connection/discuss/1295887/Easy-Solution-w-Explanation-or-All-Possible-Approaches-with-comments
package main

import "fmt"

//   1,2,3,4,5
//p [1,1,1,1,5]
//r [4,1,1,1,1]

// 2 -> 1 <-3

//navigate through the edges. [1,2] => p[1,2] -> p[1,1]
// [1,3] => r[2,1] -> r[3,1]
// [3,4] => p[1,4] r[3,1]
// [2,4] => p[1,1]

//algorithm
// get the edge
// find the parents
// if the parents are not equal -> compare ranks -> highest rank will be the parent, and the other node will be child.
// increase the rank of parent by 1. -> the child node's parent should be updated.
// go to the next edge.
// if the parent of 2 node in that edge is the same, then cycle is identified.
func main() {
	edge := findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})
	fmt.Println(edge)
}

var rank, parent []int

func findRedundantConnection(edges [][]int) []int {
	noOfNodes := len(edges) + 1
	rank = make([]int, 0, noOfNodes)
	parent = make([]int, 0, noOfNodes)
	for i := 0; i < noOfNodes; i++ {
		rank = append(rank, 1)
		parent = append(parent, i)
	}

	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		if !Union(x, y) {
			return edges[i]
		}
	}
	return []int{}
}

func findParent(n int) int {
	if parent[n] == n {
		return n
	}
	return findParent(parent[n])
}

func Union(x, y int) bool {
	a, b := findParent(x), findParent(y)
	if a == b {
		return false
	}
	if rank[x] >= rank[y] {
		rank[x] = rank[x] + 1
		parent[b] = a
		return true
	}

	rank[y] = rank[y] + 1
	parent[a] = b
	return true
}
