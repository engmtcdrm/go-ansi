package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for [Strip] function.
func Test_Strip(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\x1b[31mHello\x1b[0m", "Hello"},
		{"\x1b[1;32mBold Green\x1b[0m", "Bold Green"},
		{"\x1b[4;33mUnderlined Yellow\x1b[0m", "Underlined Yellow"},
		{"\x1b[38;5;82m256 Color Green\x1b[0m", "256 Color Green"},
		{"\x1b[48;2;255;0;0mTrue Color Red Background\x1b[0m", "True Color Red Background"},
		{"\x1b]0;Title\x07", "\x1b]0;Title\x07"}, // OSC (Operating System Command) sequence
		{"\x1b[?25lHidden Cursor\x1b[?25h", "Hidden Cursor"},
		{"\x1b[2J\x1b[HClear Screen", "Clear Screen"},
		{"\x1b[1;31;42mMultiple Attributes\x1b[0m", "Multiple Attributes"},
		{"\x1b[0mReset\x1b[0m", "Reset"},
		// Edge cases
		{"Plain text", "Plain text"}, // No ANSI codes
		{"", ""},                     // Empty string
		{"\x1b[31mRed\x1b[32mGreen\x1b[0m", "RedGreen"},     // Nested ANSI codes
		{"\x1b[31", "\x1b[31"},                              // Incomplete ANSI code
		{"\x1b[", "\x1b["},                                  // Malformed ANSI code
		{"\x1b[mText\x1b[m", "Text"},                        // Empty parameter ANSI codes
		{"Text\x1b[0m", "Text"},                             // ANSI code at end only
		{"\x1b[0mText", "Text"},                             // ANSI code at start only
		{"Multi\x1b[31mple\x1b[0m words", "Multiple words"}, // ANSI codes in middle
		{"\033[31mOctal escape\033[0m", "Octal escape"},     // Octal escape representation
	}

	for _, test := range tests {
		result := Strip(test.input)
		require.Equal(t, test.expected, result, "Strip(%q) = %q; want %q", test.input, result, test.expected)
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
	t.Run("Valid range 100", func(t *testing.T) {
		require.True(t, colorInRange(100), "colorInRange(100) should be true")
	})

	t.Run("Valid range 0", func(t *testing.T) {
		require.True(t, colorInRange(0), "colorInRange(0) should be true")
	})

	t.Run("Valid range 255", func(t *testing.T) {
		require.True(t, colorInRange(255), "colorInRange(255) should be true")
	})

	t.Run("Invalid range -1", func(t *testing.T) {
		require.False(t, colorInRange(-1), "colorInRange(-1) should be false")
	})

	t.Run("Invalid range 256", func(t *testing.T) {
		require.False(t, colorInRange(256), "colorInRange(256) should be false")
	})
}
