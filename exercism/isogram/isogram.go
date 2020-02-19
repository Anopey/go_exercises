package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram evalues whether the given string is an isogram or not.
func IsIsogram(s string) bool {
	chars := make(map[rune]int)
	s = strings.ToUpper(s)
	for _, v := range s {
		if !unicode.IsLetter(v) {
			continue
		}
		if _, ok := chars[v]; ok {
			return false
		} else {
			chars[v] = 0
		}
	}
	return true
}
