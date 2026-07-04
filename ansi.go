package ansi

const (
	Escape = "\x1b"       // ANSI escape character.
	CSI    = Escape + "[" // ANSI control sequence introducer.
)

// Strip removes all 7-bit ANSI escape sequences from the input.
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

// StripRunes removes all 7-bit ANSI escape sequences from input.
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

// Strip8Bit removes all 8-bit C1 ANSI escape sequences from the input.
func Strip8Bit[T ~string | ~[]byte](input T) T {
	var stripedString []byte

	for len(input) > 0 {
		ansiLen := escapeLength8Bit(input)

		if ansiLen == 0 {
			stripedString = append(stripedString, input[0])
			input = input[1:]
			continue
		}

		input = input[ansiLen:]
	}

	return T(stripedString)
}

// Strip8BitRunes removes all 8-bit C1 ANSI escape sequences from the input.
func Strip8BitRunes[T ~[]rune](input T) T {
	var stripedRunes []rune

	for len(input) > 0 {
		ansiLen := escapeLength8BitRune(input)

		if ansiLen == 0 {
			stripedRunes = append(stripedRunes, input[0])
			input = input[1:]
			continue
		}

		input = input[ansiLen:]
	}

	return T(stripedRunes)
}

// colorInRange checks if the color value is between 0 and 255.
func colorInRange(color int) bool {
	return color >= 0 && color <= 255
}
