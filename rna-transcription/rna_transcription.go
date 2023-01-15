package strand

var dnamap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := ""

	for _, value := range dna {
		rna = rna + string(dnamap[value])
	}

	return rna
}
