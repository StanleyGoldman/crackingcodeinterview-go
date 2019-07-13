package unique

func isUnique(value string) bool {

	m := make(map[rune]bool)

	for _, c := range value {
		if !m[c] {
			m[c] = true
		} else {
			return false
		}
	}

	return true
}

func isUnique2(value string) bool {
	m := 0

	for _, c := range value {
		v := uint(c) - uint('a')
		s := 1 << v
		if m&(s) > 0 {
			return false
		}

		m |= s
	}

	return true
}
