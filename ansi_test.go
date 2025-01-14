package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestForeground8Bit(t *testing.T) {
	tests := []struct {
		color    int
		expected string
	}{
		{0, "\x1b[38;5;0m"},
		{255, "\x1b[38;5;255m"},
		{-1, ""},
		{256, ""},
	}

	for _, test := range tests {
		result := ansi.Foreground8Bit(test.color)
		if result != test.expected {
			t.Errorf("Foreground8Bit(%d) = %q; want %q", test.color, result, test.expected)
		}
	}
}

func TestBackground8Bit(t *testing.T) {
	tests := []struct {
		color    int
		expected string
	}{
		{0, "\x1b[48;5;0m"},
		{255, "\x1b[48;5;255m"},
		{-1, ""},
		{256, ""},
	}

	for _, test := range tests {
		result := ansi.Background8Bit(test.color)
		if result != test.expected {
			t.Errorf("Background8Bit(%d) = %q; want %q", test.color, result, test.expected)
		}
	}
}

func TestForeground24Bit(t *testing.T) {
	tests := []struct {
		r, g, b  int
		expected string
	}{
		{0, 0, 0, "\x1b[38;2;0;0;0m"},
		{255, 255, 255, "\x1b[38;2;255;255;255m"},
		{-1, 0, 0, ""},
		{0, -1, 0, ""},
		{0, 0, -1, ""},
		{256, 0, 0, ""},
		{0, 256, 0, ""},
		{0, 0, 256, ""},
	}

	for _, test := range tests {
		result := ansi.Foreground24Bit(test.r, test.g, test.b)
		if result != test.expected {
			t.Errorf("Foreground24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		}
	}
}

func TestBackground24Bit(t *testing.T) {
	tests := []struct {
		r, g, b  int
		expected string
	}{
		{0, 0, 0, "\x1b[48;2;0;0;0m"},
		{255, 255, 255, "\x1b[48;2;255;255;255m"},
		{-1, 0, 0, ""},
		{0, -1, 0, ""},
		{0, 0, -1, ""},
		{256, 0, 0, ""},
		{0, 256, 0, ""},
		{0, 0, 256, ""},
	}

	for _, test := range tests {
		result := ansi.Background24Bit(test.r, test.g, test.b)
		if result != test.expected {
			t.Errorf("Background24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		}
	}
}

func TestCursorMovement(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) string
		arg      int
		expected string
	}{
		{"CursorUp", ansi.CursorUp, 1, "\x1b[1A"},
		{"CursorDown", ansi.CursorDown, 1, "\x1b[1B"},
		{"CursorForward", ansi.CursorForward, 1, "\x1b[1C"},
		{"CursorBackward", ansi.CursorBackward, 1, "\x1b[1D"},
		{"CursorNextLineN", ansi.CursorNextLineN, 1, "\x1b[1E"},
		{"CursorPreviousLineN", ansi.CursorPreviousLineN, 1, "\x1b[1F"},
		{"CursorHorizontalAbsolute", ansi.CursorHorizontalAbsolute, 1, "\x1b[1G"},
	}

	for _, test := range tests {
		result := test.function(test.arg)
		if result != test.expected {
			t.Errorf("%s(%d) = %q; want %q", test.name, test.arg, result, test.expected)
		}
	}
}

func TestCursorPosition(t *testing.T) {
	tests := []struct {
		row, column int
		expected    string
	}{
		{1, 1, "\x1b[1;1H"},
		{10, 20, "\x1b[10;20H"},
	}

	for _, test := range tests {
		result := ansi.CursorPosition(test.row, test.column)
		if result != test.expected {
			t.Errorf("CursorPosition(%d, %d) = %q; want %q", test.row, test.column, result, test.expected)
		}
	}
}

func TestScroll(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) string
		arg      int
		expected string
	}{
		{"ScrollUpN", ansi.ScrollUpN, 1, "\x1b[1S"},
		{"ScrollDownN", ansi.ScrollDownN, 1, "\x1b[1T"},
	}

	for _, test := range tests {
		result := test.function(test.arg)
		if result != test.expected {
			t.Errorf("%s(%d) = %q; want %q", test.name, test.arg, result, test.expected)
		}
	}
}

func TestStripCodes(t *testing.T) {
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
	}

	for _, test := range tests {
		result := ansi.StripCodes(test.input)
		if result != test.expected {
			t.Errorf("StripCodes(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
