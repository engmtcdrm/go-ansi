package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

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
