package hamming

type errorBoi struct {
	errorMessage string
}

func (e *errorBoi) Error() string {
	return e.errorMessage
}

//Distance calculates the amming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, &errorBoi{"noo"}
	}
	return len(a) - GetSharedCharCount(a, b), nil
}

//GetSharedCharCount returns the number of shared chars in the given two ASCII strings
func GetSharedCharCount(a, b string) int {
	max := len(a)
	if len(b) < len(a) {
		max = len(b)
	}
	returned := 0
	for i := 0; i < max; i++ {
		if a[i] == b[i] {
			returned++
		}
	}
	return returned
}
