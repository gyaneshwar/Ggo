package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("program started")
	result := sortSentence("is2 sentence4 This1 a3")
	if result == "This is a sentence" {
		fmt.Println("awesoem")
		return
	}

	fmt.Printf(result)
	fmt.Println("failure")
}

type word struct {
	name  []rune
	order int
}

func (w word) Compare(y word) bool {
	return w.order < y.order
}

//https://leetcode.com/problems/sorting-the-sentence/
func sortSentence(s string) string {
	if s == "" {
		return ""
	}
	words := strings.Split(s, " ")
	wrs := make([]word, 0, len(words))
	for _, w := range words {
		wr := word{}
		temp := []rune(w)
		wr.name = temp[:len(temp)-1]
		wr.order = int(temp[len(temp)-1])
		wrs = append(wrs, wr)
	}

	QS(wrs)

	var result []rune
	for i, v := range wrs {
		result = append(result, v.name...)
		if len(wrs)-1 != i {
			result = append(result, rune(' '))
		}
	}

	return string(result)
}

func QS(i []word) {
	if len(i) < 2 {
		return
	}

	pivot, left, right := 0, 0, len(i)-1

	i[pivot], i[right] = i[right], i[pivot]

	for index, v := range i {
		if v.Compare(i[right]) {
			i[left], i[index] = i[index], i[left]
			left++
		}
	}

	i[right], i[left] = i[left], i[right]

	QS(i[:left])
	QS(i[left+1:])
}
