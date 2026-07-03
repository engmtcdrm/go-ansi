package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type stripTestCase struct {
	name     string
	input    string
	expected string
}

var stripTests = []stripTestCase{
	{"Red text", "\x1b[31mHello\x1b[0m", "Hello"},
	{"Bold Green text", "\x1b[1;32mBold Green\x1b[0m", "Bold Green"},
	{"Underlined Yellow text", "\x1b[4;33mUnderlined Yellow\x1b[0m", "Underlined Yellow"},
	{"256 Color Green text", "\x1b[38;5;82m256 Color Green\x1b[0m", "256 Color Green"},
	{"True Color Red Background", "\x1b[48;2;255;0;0mTrue Color Red Background\x1b[0m", "True Color Red Background"},
	{"OSC (Operating System Command) sequence", "\x1b]0;Title\x07", ""},
	{"Hidden Cursor", "\x1b[?25lHidden Cursor\x1b[?25h", "Hidden Cursor"},
	{"Clear Screen", "\x1b[2J\x1b[HClear Screen", "Clear Screen"},
	{"Multiple Attributes", "\x1b[1;31;42mMultiple Attributes\x1b[0m", "Multiple Attributes"},
	{"Reset", "\x1b[0mReset\x1b[0m", "Reset"},
	{"No ANSI codes", "Plain text", "Plain text"},
	{"Empty string", "", ""},
	{"Nested ANSI codes", "\x1b[31mRed\x1b[32mGreen\x1b[0m", "RedGreen"},
	{"Incomplete ANSI code", "\x1b[31", "\x1b[31"},
	{"Malformed ANSI code", "\x1b[", "\x1b["},
	{"Empty parameter ANSI codes", "\x1b[mText\x1b[m", "Text"},
	{"ANSI code at end only", "Text\x1b[0m", "Text"},
	{"ANSI code at start only", "\x1b[0mText", "Text"},
	{"ANSI codes in middle", "Multi\x1b[31mple\x1b[0m words", "Multiple words"},
	{"Octal escape representation", "\033[31mOctal escape\033[0m", "Octal escape"},
}

// Tests for [Strip] function.
func Test_Strip(t *testing.T) {
	t.Parallel()

	var tests = make([]stripTestCase, len(stripTests))
	copy(tests, stripTests)

	type customType string

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Strip(tt.input)
			require.Equal(t, tt.expected, result, "Strip(%q) = %q; want %q", tt.input, result, tt.expected)

			resultByte := Strip([]byte(tt.input))
			require.Equal(t, tt.expected, string(resultByte), "Strip([]byte(%q)) = %q; want %q", tt.input, resultByte, tt.expected)

			resultCustom := Strip(customType(tt.input))
			require.Equal(t, customType(tt.expected), resultCustom, "Strip(customType(%q)) = %q; want %q", tt.input, resultCustom, tt.expected)
		})
	}
}

// Tests for [StripRunes] function.
func Test_StripRunes(t *testing.T) {
	t.Parallel()

	var tests = make([]stripTestCase, len(stripTests))
	copy(tests, stripTests)

	type customType2 []rune

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resultRune := StripRunes([]rune(tt.input))
			require.Equal(t, tt.expected, string(resultRune), "Strip([]rune(%q)) = %q; want %q", tt.input, resultRune, tt.expected)

			resultCustom2 := StripRunes(customType2([]rune(tt.input)))
			resultCustomExpected := customType2([]rune(tt.expected))
			require.ElementsMatch(t, resultCustomExpected, resultCustom2, "Strip(customType2(%q)) = %q; want %q", tt.input, resultCustom2, tt.expected)
		})
	}
}

// Tests for [StripCodes] function to ensure it behaves the same as Strip.
func Test_StripCodes(t *testing.T) {
	input := "\x1b[31mHello\x1b[0m"
	expected := "Hello"
	result := StripCodes(input)
	require.Equal(t, expected, result, "StripCodes(%q) = %q; want %q", input, result, expected)
}

// Tests for [colorInRange] function.
func Test_colorInRange(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Invalid range -1", -1, false},
		{"Valid range 0", 0, true},
		{"Valid range 100", 100, true},
		{"Valid range 255", 255, true},
		{"Invalid range 256", 256, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := colorInRange(tt.input)
			require.Equal(t, tt.expected, result, "colorInRange(%d) = %v; want %v", tt.input, result, tt.expected)
		})
	}
}

func Benchmark_Strip(b *testing.B) {
	input := "\x1b[31mHello\x1b[0m"
	for i := 0; i < b.N; i++ {
		_ = Strip(input)
	}
}
