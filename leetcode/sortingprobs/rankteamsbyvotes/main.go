package main

import "fmt"

//https://leetcode.com/problems/rank-teams-by-votes/
func main() {
	if rankTeams([]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}) == "ABC" {
		fmt.Print("awesome")
		return
	}
	fmt.Print("failed")
}

var alphabets string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Helpers
func findPosition(str string, alphabet string) int {
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		if alphabet == string(chars[i]) {
			return i
		}
	}
	return 0
}

type Alphabet struct {
	item              string //A
	alphabeticalOrder int    //1
	positionalvotes   []int  //[5,0,0]
}

// create a collection of A-Z using alphabets and its positions
// collect the positions of alphabets in each set.
// while collecting the positions of an alphabet in a particular set, increment the votes by one in its respective position.

//sorting
// sort alphabets by votes
// 			compare function between two alphabets
//			if count of a position of an alphabet is less than count of position of another alphabet, throw all of them to left (A.positionalvotes[0] < B.positionalvotes[0])
//			if the above thing doesnt render a result, then order them by alphabeticalOrder.
func rankTeams(votes []string) string {
	if len(votes) == 0 {
		return ""
	}
	// if len less than 2, return
	if len(votes) == 1 {
		return votes[0]
	}

	//get the first element and figure out the alphabets that we need to deal with.
	alp := []Alphabet{}

	// initialize
	firstValue := []rune(votes[0])

	for i := 0; i < len(firstValue); i++ {
		val := string(firstValue[i])
		alp = append(alp, Alphabet{item: val, alphabeticalOrder: findPosition(alphabets, val), positionalvotes: make([]int, len(firstValue))})
	}

	for i := 0; i < len(votes); i++ {
		for index, a := range alp {
			p := findPosition(votes[i], a.item)
			alp[index].positionalvotes[p]++
		}
	}

	QS(alp)

	var result string
	for _, v := range alp {
		result = fmt.Sprintf("%s%s", result, v.item)
	}

	return result

}

func QS(alpabets []Alphabet) {
	if len(alpabets) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(alpabets)-1

	alpabets[pivot], alpabets[right] = alpabets[right], alpabets[pivot]

	for i := range alpabets {
		if xgreaterthany(alpabets[i], alpabets[right]) {
			alpabets[i], alpabets[left] = alpabets[left], alpabets[i]
			left++
		}
	}

	alpabets[right], alpabets[left] = alpabets[left], alpabets[right]

	QS(alpabets[:left])
	QS(alpabets[left+1:])
}

func xgreaterthany(a Alphabet, b Alphabet) bool {
	for i := 0; i < len(a.positionalvotes); i++ {
		if a.positionalvotes[i] > b.positionalvotes[i] {
			return true
		}
		if a.positionalvotes[i] < b.positionalvotes[i] {
			return false
		}
	}

	// alphabet position
	return a.alphabeticalOrder < b.alphabeticalOrder
}
