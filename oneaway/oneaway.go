package oneaway

func oneAway(x string, y string) bool {
	results := make(map[rune]int)

	for _, c := range x {
		results[c]++
	}

	for _, c := range y {
		results[c]--
	}

	left := 0
	right := 0
	for _, v := range results {
		if v > 0 {
			left += v
		} else if v < 0 {
			right += v
		}
	}

	return (left <= 1) && (right <= 1)
}
