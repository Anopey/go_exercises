package raindrops

import "strconv"

//Convert converts a number to raindrop sounds :o
func Convert(num int) string {
	returned := ""
	if num%3 == 0 {
		returned += "Pling"
	}
	if num%5 == 0 {
		returned += "Plang"
	}
	if num%7 == 0 {
		returned += "Plong"
	}
	if len(returned) == 0 {
		return strconv.Itoa(num)
	}
	return returned
}
