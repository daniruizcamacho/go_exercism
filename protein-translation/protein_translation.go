package protein

import (
	"errors"
)

var codonsTable = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}
var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

func FromCodon(codon string) (string, error) {
	aminoacid, ok := codonsTable[codon]
	if !ok {
		return "", ErrInvalidBase
	} else if aminoacid == "STOP" {
		return "", ErrStop
	} else {
		return aminoacid, nil
	}
}
func FromRNA(rna string) ([]string, error) {
	var result []string
	for nt := 0; nt <= len(rna)-3; nt += 3 {
		aminoacid, ok := FromCodon(rna[nt : nt+3])
		if ok == ErrInvalidBase {
			return result, ErrInvalidBase
		} else if ok == ErrStop {
			return result, nil
		}
		result = append(result, aminoacid)
	}
	return result, nil
}
