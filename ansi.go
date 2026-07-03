package ansi

import (
	"regexp"
)

const (
	Escape = "\x1b"       // ANSI escape character.
	CSI    = Escape + "[" // ANSI control sequence introducer.
)

var ansiRegex = regexp.MustCompile(`\x1b\[[;?0-9]*[a-zA-Z]`)

// Strip removes all ANSI escape codes from the input.
func Strip[T ~string | ~[]byte | ~[]rune](input T) T {
	switch v := any(input).(type) {
	case string:
		return T(ansiRegex.ReplaceAllString(v, ""))
	default:
		return T(ansiRegex.ReplaceAllString(string(input), ""))
	}
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
