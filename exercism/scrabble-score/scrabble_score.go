package scrabble

import "strings"

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10}

//Score returns the scrabble score of a given string
func Score(s string) int {
	s = strings.ToUpper(s)
	score := 0
	for _, v := range s {
		for mk, mv := range scores {
			if strings.ContainsRune(mk, v) {
				score += mv
				break
			}
		}
	}
	return score
}
