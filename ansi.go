package ansi

import (
	"regexp"
)

const (
	Escape = "\x1b"       // ANSI escape character
	CSI    = Escape + "[" // ANSI control sequence introducer
)

var ansiRegex = regexp.MustCompile(`(?:\x1b|\033)\[[;?0-9]*[a-zA-Z]`)

// Strip removes all ANSI escape codes from the input string.
func Strip(input string) string {
	return ansiRegex.ReplaceAllString(input, "")
}

// StripCodes removes all ANSI escape codes from the input string.
//
// Deprecated: Use [Strip] instead. This will be removed in v2.0.0.
func StripCodes(input string) string {
	return Strip(input)
}

// colorInRange checks if the color value is between 0 and 255.
func colorInRange(color int) bool {
	return color >= 0 && color <= 255
}
