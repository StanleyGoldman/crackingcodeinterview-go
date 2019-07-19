package compression

import (
	"strconv"
	"strings"
)

func compressString(input string) string {
	var sb strings.Builder

	var b rune
	var bCount int
	for i, c := range input {
		if i > 0 {
			if c == b {
				bCount++
				continue
			} else {
				sb.WriteRune(b)
				sb.WriteString(strconv.Itoa(bCount))
			}
		}

		b = c
		bCount = 1
	}

	sb.WriteRune(b)
	sb.WriteString(strconv.Itoa(bCount))

	output := sb.String()

	if len(input) <= len(output) {
		return input
	}

	return output
}
