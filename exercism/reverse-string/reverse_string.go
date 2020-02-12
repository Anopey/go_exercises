//Package reverse implements the solution to the string reversal problem.
package reverse

import "unicode/utf8"

//Reverse function reverses a given string
func Reverse(inp string) string {
	var output string
	for i := utf8.RuneCountInString(inp) - 1; i >= 0; i-- {
		output += string(inp[i])
	}
	return output
}
