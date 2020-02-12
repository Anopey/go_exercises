//Package reverse implements the solution to the string reversal problem.
package reverse

//Reverse function reverses a given string
func Reverse(inp string) string {
	var output string
	runes := []rune(inp)
	for i := len(runes) - 1; i >= 0; i-- {
		output += string(runes[i])
	}
	return output
}
