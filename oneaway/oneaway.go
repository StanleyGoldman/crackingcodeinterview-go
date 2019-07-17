package oneaway

import "math"

func oneAway(x string, y string) bool {
	xBig := len(x) > len(y)
	yBig := len(y) > len(x)
	same := len(x) == len(y)

	if !same {
		if int(math.Abs(float64(len(x)-len(y)))) >= 2 {
			return false
		}
	}

	p := int(math.Min(float64(len(x)), float64(len(y))))
	xi := 0
	yi := 0
	diff := 0

	for i := 0; i < p; i++ {
		if xi > len(x) || yi > len(y) {
			diff++
		} else if x[xi] == y[yi] {
			xi++
			yi++
		} else {
			diff++

			if same {
				xi++
				yi++
			} else if xBig {
				xi++
			} else if yBig {
				yi++
			}
		}

		if diff > 1 {
			return false
		}
	}

	return true
}
