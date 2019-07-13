package urlify

func urlify(i string, l int) string {
	ir := []rune(i)
	iri := len(ir) - 1

	for i := l - 1; i > 0 && i <= iri; i-- {
		if ir[i] == ' ' {
			ir[iri] = '0'
			iri--
			ir[iri] = '2'
			iri--
			ir[iri] = '%'
			iri--
		} else {
			ir[iri] = ir[i]
			iri--
		}
	}

	return string(ir)
}
