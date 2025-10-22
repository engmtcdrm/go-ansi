package ansi

import "strconv"

var (
	ScrollUp1   = ScrollUpN(1)   // Scroll the screen up one row (line).
	ScrollDown1 = ScrollDownN(1) // Scroll the screen down one row (line).
)

// ScrollUp scrolls the screen up n rows (lines).
// If n is negative, it returns an empty string.
func ScrollUpN(n int) string {
	if n < 0 {
		return ""
	}
	return csi + strconv.Itoa(n) + "S"
}

// ScrollDownN scrolls the screen down n rows (lines).
// If n is negative, it returns an empty string.
func ScrollDownN(n int) string {
	if n < 0 {
		return ""
	}
	return csi + strconv.Itoa(n) + "T"
}
