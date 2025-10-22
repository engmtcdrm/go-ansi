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
		// Edge cases - negative values
		{"CursorUp negative", ansi.CursorUp, -1, ""},
		{"CursorDown negative", ansi.CursorDown, -1, ""},
		{"CursorForward negative", ansi.CursorForward, -1, ""},
		{"CursorBackward negative", ansi.CursorBackward, -1, ""},
		{"CursorNextLineN negative", ansi.CursorNextLineN, -1, ""},
		{"CursorPreviousLineN negative", ansi.CursorPreviousLineN, -1, ""},
		{"CursorHorizontalAbsolute negative", ansi.CursorHorizontalAbsolute, -1, ""},
		// Edge cases - zero values
		{"CursorUp zero", ansi.CursorUp, 0, ""},
		{"CursorDown zero", ansi.CursorDown, 0, ""},
		{"CursorForward zero", ansi.CursorForward, 0, "\x1b[0C"},
		{"CursorBackward zero", ansi.CursorBackward, 0, "\x1b[0D"},
		{"CursorNextLineN zero", ansi.CursorNextLineN, 0, "\x1b[0E"},
		{"CursorPreviousLineN zero", ansi.CursorPreviousLineN, 0, "\x1b[0F"},
		{"CursorHorizontalAbsolute zero", ansi.CursorHorizontalAbsolute, 0, ""},
		// Large values
		{"CursorUp large", ansi.CursorUp, 100, "\x1b[100A"},
		{"CursorDown large", ansi.CursorDown, 999, "\x1b[999B"},
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
		// Edge cases
		{0, 0, ""},                  // Both zero
		{-1, -1, ""},                // Both negative
		{0, 1, ""},                  // Row zero
		{1, 0, ""},                  // Column zero
		{-1, 1, ""},                 // Row negative
		{1, -1, ""},                 // Column negative
		{100, 200, "\x1b[100;200H"}, // Large values
	}

	for _, test := range tests {
		result := ansi.CursorPosition(test.row, test.column)
		if result != test.expected {
			t.Errorf("CursorPosition(%d, %d) = %q; want %q", test.row, test.column, result, test.expected)
		}
	}
}

func TestCursorDerivedConstants(t *testing.T) {
	// Test composite/derived constants
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{"CursorTopLeft", ansi.CursorTopLeft, ansi.CursorPosition(1, 1)},
		{"CursorLineBegin", ansi.CursorLineBegin, ansi.CursorHorizontalAbsolute(1)},
		{"CursorNextLine", ansi.CursorNextLine, ansi.CursorNextLineN(1)},
		{"CursorPreviousLine", ansi.CursorPreviousLine, ansi.CursorPreviousLineN(1)},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("%s = %q; want %q", test.name, test.constant, test.expected)
		}
	}
}
