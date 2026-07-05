package ansi

import "strconv"

var (
	ScrollUp1   = ScrollUp(1)   // Scroll the screen up one row (line).
	ScrollDown1 = ScrollDown(1) // Scroll the screen down one row (line).
)

// ScrollUp scrolls the screen up n rows (lines).
// If n is less than 1, it returns an empty string.
func ScrollUp(n int) string {
	if n < 1 {
		return ""
	}
	return CSI + strconv.Itoa(n) + "S"
}

// ScrollDown scrolls the screen down n rows (lines).
// If n is less than 1, it returns an empty string.
func ScrollDown(n int) string {
	if n < 1 {
		return ""
	}
	return CSI + strconv.Itoa(n) + "T"
}
