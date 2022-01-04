package main

func main() {}

func reverseShuffleMerge(s string) (end string) {

	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		end = end + string(r[i])
	}

	return end
}
