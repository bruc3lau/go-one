package main

import "fmt"

func main() {
	//fmt.Println(666)
	result := solve("abcab")
	fmt.Println(result)
}

func solve(s string) int {
	//fmt.Println(input)
	//inputLen := len(input)
	//fmt.Println(inputLen)
	//for a := range input {
	//	fmt.Println(string(input[a]))
	//}

	//return input

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for end := 0; end < len(s); end++ {
		if lastI, ok := lastOccurred[s[end]]; ok && lastI >= start {
			start = lastI + 1
		}
		lastOccurred[s[end]] = end

		curLength := end - start + 1
		if curLength > maxLength {
			maxLength = curLength
		}
	}

	return maxLength
}
