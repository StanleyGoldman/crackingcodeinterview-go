package palinpermut

import "strings"

func isPalindromePermutation(input string) bool {
	input = strings.ToLower(input)
	count := 0
	runecounts := make(map[rune]int)
	inputrunes := []rune(input)

	for _, c := range inputrunes {
		if c != ' ' {
			count++
			runecounts[c]++
		}
	}

	if count%2 == 0 {
		for _, c := range runecounts {
			if c%2 != 0 {
				return false
			}
		}
	} else {
		oneOdd := false
		for _, c := range runecounts {
			if c%2 != 0 {
				if oneOdd {
					return false
				}

				oneOdd = true
			}
		}
	}

	return true
}
