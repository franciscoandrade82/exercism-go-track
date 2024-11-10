package strand

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna string
	for _, letter := range dna {
		rna += string(dnaToRna[letter])
	}

	return rna
}
