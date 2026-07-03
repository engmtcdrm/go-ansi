package ansi

const (
	Escape = "\x1b"       // ANSI escape character.
	CSI    = Escape + "[" // ANSI control sequence introducer.
)

// Strip removes all ANSI escape sequences from the input.
func Strip[T ~string | ~[]byte](input T) T {
	var stripedString []byte

	for len(input) > 0 {
		ansiLen := escapeLength(input)

		if ansiLen == 0 {
			stripedString = append(stripedString, input[0])
			input = input[1:]
			continue
		}

		input = input[ansiLen:]
	}

	return T(stripedString)
}

// StripRunes removes all ANSI escape sequences from a slice of runes.
func StripRunes[T ~[]rune](input T) T {
	var stripedRunes []rune

	for len(input) > 0 {
		ansiLen := escapeLengthRune(input)

		if ansiLen == 0 {
			stripedRunes = append(stripedRunes, input[0])
			input = input[1:]
			continue
		}

		input = input[ansiLen:]
	}

	return T(stripedRunes)
}

// StripCodes removes all ANSI escape sequences from the input string.
//
// Deprecated: Use [Strip] instead. This will be removed in v2.0.0.
func StripCodes(input string) string {
	return Strip(input)
}

// colorInRange checks if the color value is between 0 and 255.
func colorInRange(color int) bool {
	return color >= 0 && color <= 255
}
