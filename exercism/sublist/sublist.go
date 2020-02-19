package sublist

//Sublist evaluates the sublist relation between A and B.
func Sublist(A, B []int) string {
	ab := false
	ba := false
	if isSublist(A, B) {
		ab = true
	}
	if isSublist(B, A) {
		ba = true
	}

	if ab && ba {
		return "equal"
	} else if ab {
		return "sublist"
	} else if ba {
		return "superlist"
	} else {
		return "unequal"
	}
}

func isSublist(A, B []int) bool {
	if len(A) == 0 {
		return true
	}
	for i, j := 0, 0; i < len(B); i++ {
		if B[i] == A[j] {
			j++
		} else if j > 0 {
			i -= j
			j = 0
		}
		if j == len(A) {
			return true
		}
	}
	return false
}
