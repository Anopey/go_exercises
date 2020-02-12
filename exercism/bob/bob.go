//Package bob is a simulation of how bob communicates.
package bob

import "unicode"

//Hey returns the answer bob would give to a remark.
func Hey(remark string) string {
	var isShout bool = true
	var hasLetters bool = false
	var isSilence bool = true
	var isQuestion bool = false
	for _, v := range remark {
		if v == ' ' || v == '\t' || v == '\n' || v == '\r' {
			continue
		}
		isSilence = false
		isQuestion = false
		if unicode.IsLetter(rune(v)) {
			hasLetters = true
			if !unicode.IsUpper(rune(v)) {
				isShout = false
			}
		} else if v == '?' {
			isQuestion = true
		}
	}
	if isSilence {
		return "Fine. Be that way!"
	}
	if !hasLetters {
		isShout = false
	}
	if isShout {
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	return "Whatever."
}
