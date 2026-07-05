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

type customStringType string
type customBytesType []byte
type customRunesType []rune

var stripTests = []stripTestCase{
	{name: "Red text", input: "\x1b[31mHello\x1b[0m", expected: "Hello"},
	{name: "Bold Green text", input: "\x1b[1;32mBold Green\x1b[0m", expected: "Bold Green"},
	{name: "Underlined Yellow text", input: "\x1b[4;33mUnderlined Yellow\x1b[0m", expected: "Underlined Yellow"},
	{name: "256 Color Green text", input: "\x1b[38;5;82m256 Color Green\x1b[0m", expected: "256 Color Green"},
	{name: "True Color Red Background", input: "\x1b[48;2;255;0;0mTrue Color Red Background\x1b[0m", expected: "True Color Red Background"},
	{name: "OSC (Operating System Command) sequence", input: "\x1b]0;Title\x07", expected: ""},
	{name: "Hidden Cursor", input: "\x1b[?25lHidden Cursor\x1b[?25h", expected: "Hidden Cursor"},
	{name: "Clear Screen", input: "\x1b[2J\x1b[HClear Screen", expected: "Clear Screen"},
	{name: "Multiple Attributes", input: "\x1b[1;31;42mMultiple Attributes\x1b[0m", expected: "Multiple Attributes"},
	{name: "Reset", input: "\x1b[0mReset\x1b[0m", expected: "Reset"},
	{name: "No ANSI codes", input: "Plain text", expected: "Plain text"},
	{name: "Empty string", input: "", expected: ""},
	{name: "Nested ANSI codes", input: "\x1b[31mRed\x1b[32mGreen\x1b[0m", expected: "RedGreen"},
	{name: "Incomplete ANSI code", input: "\x1b[31", expected: "\x1b[31"},
	{name: "Malformed ANSI code", input: "\x1b[", expected: "\x1b["},
	{name: "Empty parameter ANSI codes", input: "\x1b[mText\x1b[m", expected: "Text"},
	{name: "ANSI code at end only", input: "Text\x1b[0m", expected: "Text"},
	{name: "ANSI code at start only", input: "\x1b[0mText", expected: "Text"},
	{name: "ANSI codes in middle", input: "Multi\x1b[31mple\x1b[0m words", expected: "Multiple words"},
	{name: "Octal escape representation", input: "\033[31mOctal escape\033[0m", expected: "Octal escape"},
}

var strip8BitTests = []stripTestCase{
	{name: "No ANSI codes", input: "Plain text", expected: "Plain text"},
	{name: "Empty string", input: "", expected: ""},
	{name: "C1 CSI then text", input: "\x9B31mhello", expected: "hello"},
}

// Benchmark for [Strip] function.
func Benchmark_Strip(b *testing.B) {
	input := "\x1b[31mHello\x1b[0m"
	for i := 0; i < b.N; i++ {
		_ = Strip(input)
	}
}

// Tests for [Strip] function.
func Test_Strip(t *testing.T) {
	t.Parallel()

	errFormat := "Strip(%q) returned incorrect result for type '%s'"
	var tests = make([]stripTestCase, len(stripTests))
	copy(tests, stripTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			result := Strip(tt.input)
			require.Equal(t, tt.expected, result, errFormat, tt.input, "string")
			resultCustomString := Strip(customStringType(tt.input))
			require.Equal(t, customStringType(tt.expected), resultCustomString, errFormat, tt.input, "customStringType")

			// []byte
			resultByte := Strip([]byte(tt.input))
			require.Equal(t, tt.expected, string(resultByte), errFormat, tt.input, "[]byte")
			resultCustomBytes := Strip(customBytesType([]byte(tt.input)))
			require.ElementsMatch(t, customBytesType([]byte(tt.expected)), resultCustomBytes, errFormat, tt.input, "customBytesType")
		})
	}
}

// Benchmark for [StripRunes] function.
func Benchmark_StripRunes(b *testing.B) {
	input := []rune("\x1b[31mHello\x1b[0m")
	for i := 0; i < b.N; i++ {
		_ = StripRunes(input)
	}
}

// Tests for [StripRunes] function.
func Test_StripRunes(t *testing.T) {
	t.Parallel()

	errFormat := "StripRunes(%q) returned incorrect result for type '%T'"
	var tests = make([]stripTestCase, len(stripTests))
	copy(tests, stripTests)

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input)
			expectedRunes := stringToRunes(t, tt.expected)

			require.Equal(t, expectedRunes, StripRunes(inputRunes), errFormat, tt.input, []rune(nil))
			require.ElementsMatch(t, customRunesType(expectedRunes), StripRunes(customRunesType(inputRunes)), errFormat, tt.input, customRunesType([]rune(nil)))
		})
	}
}

// Benchmark for [Strip8Bit] function.
func Benchmark_Strip8Bit(b *testing.B) {
	input := "\x1b[31mHello\x1b[0m"
	for i := 0; i < b.N; i++ {
		_ = Strip8Bit(input)
	}
}

// Tests for [Strip8Bit] function.
func Test_Strip8Bit(t *testing.T) {
	t.Parallel()

	var tests = make([]stripTestCase, len(strip8BitTests))
	copy(tests, strip8BitTests)

	type customType string

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Strip8Bit(tt.input)
			require.Equal(t, tt.expected, result, "Strip8Bit(%q) = %q; want %q", tt.input, result, tt.expected)

			resultByte := Strip8Bit([]byte(tt.input))
			require.Equal(t, tt.expected, string(resultByte), "Strip8Bit([]byte(%q)) = %q; want %q", tt.input, resultByte, tt.expected)

			resultCustom := Strip8Bit(customType(tt.input))
			require.Equal(t, customType(tt.expected), resultCustom, "Strip8Bit(customType(%q)) = %q; want %q", tt.input, resultCustom, tt.expected)
		})
	}
}

// Benchmark for [Strip8BitRunes] function.
func Benchmark_Strip8BitRunes(b *testing.B) {
	input := []rune("\x1b[31mHello\x1b[0m")
	for i := 0; i < b.N; i++ {
		_ = Strip8BitRunes(input)
	}
}

// Tests for [Strip8BitRunes] function.
func Test_Strip8BitRunes(t *testing.T) {
	t.Parallel()

	var tests = make([]stripTestCase, len(strip8BitTests))
	copy(tests, strip8BitTests)

	type customType []rune

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input)
			resultRune := Strip8BitRunes(inputRunes)
			require.Equal(t, tt.expected, string(resultRune), "Strip([]rune(%q)) = %q; want %q", tt.input, resultRune, tt.expected)

			resultCustom2 := Strip8BitRunes(customType(inputRunes))
			resultCustomExpected := customType([]rune(tt.expected))
			require.ElementsMatch(t, resultCustomExpected, resultCustom2, "Strip(customType2(%q)) = %q; want %q", tt.input, resultCustom2, tt.expected)
		})
	}
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
