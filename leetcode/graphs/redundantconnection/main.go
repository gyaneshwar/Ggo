// https://leetcode.com/problems/redundant-connection/ - 8:26 , submitted 8:50 ,
package main

var nodeChildren map[int]map[int]bool

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	nodeChildren := make(map[int]map[int]bool)

	fillNodes := func(p, c int) {
		if children, ok := nodeChildren[p]; ok {
			children[c] = true
			nodeChildren[p] = children
			return
		}
		children := map[int]bool{c: true}
		nodeChildren[p] = children
	}
	// fill in node and its children.
	for _, edge := range edges {
		fillNodes(edge[0], edge[1])
		fillNodes(edge[1], edge[0])
	}

	// consider all the nodes which are having children greater than 1
	//traverse from end of the edges and figure out the first combination.
	for i := len(edges) - 1; i > 0; i-- {
		edge := edges[i]
		e0c := nodeChildren[edge[0]]
		e1c := nodeChildren[edge[1]]
		if len(e0c) > 1 && len(e1c) > 1 {
			return edge
		}
	}
	return []int{}
}
