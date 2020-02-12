// Package twofer contains the solution.
package twofer

import "strings"

// ShareWith returns the two-fer string
func ShareWith(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
