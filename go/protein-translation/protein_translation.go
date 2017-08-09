package protein

const testVersion = 1

var (
	codonMap = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
	}
)

const (
	stopCodon = "STOP"
)

func FromCodon(codon string) string {
	if out, ok := codonMap[codon]; ok {
		return out
	}
	return stopCodon
}

func FromRNA(rna string) []string {
	seq := []string{}
	for i := 0; i < len(rna)-2; i += 3 {
		codon := FromCodon(rna[i : i+3])
		if codon == stopCodon {
			return seq
		}
		seq = append(seq, codon)
	}
	return seq
}
