//https://leetcode.com/problems/maximal-network-rank/ 7:23 , 7:35- 8:25 - 8:30

package main

import "fmt"

func main() {
	// fmt.Println(maximalNetworkRank(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(maximalNetworkRank(5, [][]int{{2, 3}, {3, 0}, {0, 4}, {4, 1}}))
}

type city struct {
	cityID          int
	roads           int
	connectedCities map[int]bool
}

func maximalNetworkRank(n int, roads [][]int) int {
	if n == 2 {
		if len(roads) == 0 {
			return 0
		}
		return 1
	}
	// find the counter of every node.
	cityRoads := make(map[int]city)

	fillData := func(city1, city2 int) {
		if v, ok := cityRoads[city1]; !ok {
			c := city{cityID: city1, roads: 1, connectedCities: map[int]bool{city2: true}}
			cityRoads[city1] = c
		} else {
			v.roads++
			v.connectedCities[city2] = true
			cityRoads[city1] = v
		}
	}

	for _, road := range roads {
		city1 := road[0]
		city2 := road[1]

		fillData(city1, city2)
		fillData(city2, city1)
	}

	cities := make([]city, 0, len(cityRoads))
	for _, c := range cityRoads {
		cities = append(cities, c)
	}

	// order by road count and get top 2 ranks
	QS(cities)

	// pick top 2 ranks
	rank1Cities := make([]city, 0, len(cityRoads))
	maxPicks := 2
	temp := 0
	for i := 0; i < len(cities); i++ {
		//capture top 1
		if i == 0 {
			temp = cities[i].roads
			rank1Cities = append(rank1Cities, cities[i])
			maxPicks--
			continue
		}

		if temp == cities[i].roads && maxPicks == 1 {
			rank1Cities = append(rank1Cities, cities[i])
			continue
		}

		// capture top 2
		if temp != cities[i].roads && maxPicks == 1 {
			rank1Cities = append(rank1Cities, cities[i])
			temp = cities[i].roads
			maxPicks--
			continue
		}

		if temp == cities[i].roads && maxPicks == 0 {
			rank1Cities = append(rank1Cities, cities[i])
			continue
		}

		// break after capturing all top 2
		if temp != cities[i].roads {
			break
		}

	}

	var max int
	// add those 2 ranks
	// check if there is a common connection between those 2 ranks. if so negate by 1.
	for i := 0; i < len(rank1Cities); i++ {
		for j := i + 1; j < len(rank1Cities); j++ {
			total := rank1Cities[i].roads + rank1Cities[j].roads
			// find if i and j are connected
			if _, ok := cityRoads[rank1Cities[i].cityID].connectedCities[rank1Cities[j].cityID]; ok {
				total--
			}
			if total > max {
				max = total
			}
		}
	}

	return max
}

func QS(c []city) {
	if len(c) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(c)-1

	c[pivot], c[right] = c[right], c[pivot]

	for i, x := range c {
		if x.roads > c[right].roads {
			c[left], c[i] = c[i], c[left]
			left++
		}
	}

	c[right], c[left] = c[left], c[right]
	QS(c[:left])
	QS(c[left+1:])
}
