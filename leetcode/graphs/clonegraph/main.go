package main

import "fmt"

func main() {
	node1 := Node{Val: 1, Neighbors: []*Node{}}
	node2 := Node{Val: 2, Neighbors: []*Node{}}
	node3 := Node{Val: 3, Neighbors: []*Node{}}
	node4 := Node{Val: 4, Neighbors: []*Node{}}
	node1.Neighbors = append(node1.Neighbors, &node2, &node4)
	node2.Neighbors = append(node2.Neighbors, &node1, &node3)
	node3.Neighbors = append(node3.Neighbors, &node2, &node4)
	node4.Neighbors = append(node4.Neighbors, &node1, &node3)
	cloneNode1 := cloneGraph(&node1)
	fmt.Println(cloneNode1)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

var originalandclone map[*Node]*Node

func cloneGraph(original *Node) *Node {
	//traverse through the graph
	//make a copy of each node and add the child nodes to each node.
	if originalandclone == nil {
		originalandclone = make(map[*Node]*Node)
	}

	if clone, ok := originalandclone[original]; ok {
		return clone
	}
	clone := Node{Val: original.Val}
	originalandclone[original] = &clone
	for _, origChildren := range original.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneGraph(origChildren))
	}
	return &clone
}
