package main

import "fmt"

func main() {
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}

type person struct {
	townjudge bool
	childer   map[int]bool
}

func findJudge(n int, trust [][]int) int {
	if n < 1 {
		return -1
	}

	if n == 1 {
		return 1
	}

	graph := make(map[int]person)

	for _, p := range trust {
		per := p[0]
		tj := p[1]
		if v, ok := graph[tj]; ok {
			v.childer[per] = true
			graph[tj] = v
		} else {
			graph[tj] = person{townjudge: true, childer: map[int]bool{per: true}}
		}

		if v, ok := graph[per]; !ok {
			graph[per] = person{townjudge: false, childer: map[int]bool{}}
		} else {
			v.townjudge = false
			graph[per] = v
		}
	}

	tjs := make(map[int]bool)
	for per, p := range graph {
		if len(p.childer) == n-1 && p.townjudge {
			tjs[per] = true
		}
	}

	if len(tjs) > 1 {
		return -1
	}

	for tj := range tjs {
		return tj
	}

	return -1
}
