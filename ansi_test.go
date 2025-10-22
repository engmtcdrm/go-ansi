package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestStrip(t *testing.T) {
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
		result := ansi.Strip(test.input)
		if result != test.expected {
			t.Errorf("Strip(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestStripCodes(t *testing.T) {
	// Test deprecated function
	input := "\x1b[31mHello\x1b[0m"
	expected := "Hello"
	result := ansi.StripCodes(input)
	if result != expected {
		t.Errorf("StripCodes(%q) = %q; want %q", input, result, expected)
	}
}
