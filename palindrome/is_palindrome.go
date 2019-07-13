package palindrome

import "sort"

func isPalindrome(n string, h string) bool {
	if len(n) != len(h) {
		return false
	}

	slots := make([]bool, len(h))

	for _, n1 := range n {
		matched := false

		for i, h1 := range h {
			if n1 == h1 {
				if slots[i] {
					return false
				}

				slots[i] = true
				matched = true
				break
			}
		}

		if !matched {
			return false
		}
	}

	return true
}

func isPalindrome2(n string, h string) bool {
	if len(n) != len(h) {
		return false
	}

	letters := make(map[rune]int)

	for _, c := range h {
		letters[c]++
	}

	for _, c := range n {
		letters[c]--
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}

	return true
}

type RuneSlice []rune

func isPalindrome3(n string, h string) bool {
	if len(n) != len(h) {
		return false
	}

	nr := []rune(n)
	sort.Sort(RuneSlice(nr))

	hr := []rune(h)
	sort.Sort(RuneSlice(hr))

	for i, nc := range nr {
		if nc != hr[i] {
			return false
		}
	}

	return true
}

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
