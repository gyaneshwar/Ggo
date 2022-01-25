package main

import "fmt"

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
}

type nodeChildren struct {
	visited  bool
	children map[int]bool
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if n == 0 {
		return false
	}
	if n == 1 && len(edges) == 0 {
		return true
	}
	nodes := make(map[int]nodeChildren)
	for _, edge := range edges {
		n1 := edge[0]
		n2 := edge[1]

		//setting node 1
		if v, ok := nodes[n1]; !ok {
			local := make(map[int]bool)
			local[n2] = true
			nodes[n1] = nodeChildren{children: local}
		} else {
			v.children[n2] = true
			nodes[n1] = v
		}

		//setting node 2
		if v, ok := nodes[n2]; !ok {
			local := make(map[int]bool)
			local[n1] = true
			nodes[n2] = nodeChildren{children: local}
		} else {
			v.children[n1] = true
			nodes[n2] = v
		}
	}

	return search(nodes, source, destination)

}

func search(nodes map[int]nodeChildren, source, destination int) bool {
	if node, ok := nodes[source]; ok {
		if node.visited {
			return false
		}
		//assigning as visted (to avoid 1 <-> 1 )
		node.visited = true
		nodes[source] = node

		//find destination in children
		if _, ok := node.children[destination]; ok {
			return true
		}
		// try to parse the children
		for child, _ := range node.children {
			found := search(nodes, child, destination)
			if found {
				return found
			}
		}
	}
	return false
}
