package grains

import "errors"

//Square takes two to the power of i
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("integer square out of bounds or not able to be squared.")
	}
	return uint64(1) << (i - 1), nil
}

//Total takes total of grains.
func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64; i++ {
		v, e := Square(i)
		if e != nil {
			panic("Square function stopped working.")
		}
		total += v
	}
	return total
}
