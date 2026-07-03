package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
	"github.com/stretchr/testify/require"
)

type cursorTestCase struct {
	name     string
	arg      int
	expected string
}

// Tests for [ansi.CursorTopLeft] variable.
func Test_CursorTopLeft(t *testing.T) {
	t.Parallel()

	expected := ansi.CursorPosition(1, 1)
	require.Equal(t, expected, ansi.CursorTopLeft, "CursorTopLeft should be equal to its expected value.")
}

// Tests for [ansi.CursorLineBegin] variable.
func Test_CursorLineBegin(t *testing.T) {
	t.Parallel()

	expected := ansi.CursorHorizontalAbsolute(1)
	require.Equal(t, expected, ansi.CursorLineBegin, "CursorLineBegin should be equal to its expected value.")
}

// Tests for [ansi.CursorNextLine] variable.
func Test_CursorNextLine(t *testing.T) {
	t.Parallel()

	expected := ansi.CursorNextLineN(1)
	require.Equal(t, expected, ansi.CursorNextLine, "CursorNextLine should be equal to its expected value.")
}

// Tests for [ansi.CursorPreviousLine] variable.
func Test_CursorPreviousLine(t *testing.T) {
	t.Parallel()

	expected := ansi.CursorPreviousLineN(1)
	require.Equal(t, expected, ansi.CursorPreviousLine, "CursorPreviousLine should be equal to its expected value.")
}

// Tests for [ansi.CursorUp] function.
func Test_CursorUp(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorUp one", arg: 1, expected: "\x1b[A"},
		{name: "CursorUp negative", arg: -1, expected: ""},
		{name: "CursorUp zero", arg: 0, expected: ""},
		{name: "CursorUp large", arg: 100, expected: "\x1b[100A"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorUp(tt.arg)
			require.Equal(t, tt.expected, result, "CursorUp(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorDown] function.
func Test_CursorDown(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorDown one", arg: 1, expected: "\x1b[B"},
		{name: "CursorDown negative", arg: -1, expected: ""},
		{name: "CursorDown zero", arg: 0, expected: ""},
		{name: "CursorDown large", arg: 100, expected: "\x1b[100B"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorDown(tt.arg)
			require.Equal(t, tt.expected, result, "CursorDown(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorForward] function.
func Test_CursorForward(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorForward one", arg: 1, expected: "\x1b[C"},
		{name: "CursorForward negative", arg: -1, expected: ""},
		{name: "CursorForward zero", arg: 0, expected: ""},
		{name: "CursorForward large", arg: 100, expected: "\x1b[100C"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorForward(tt.arg)
			require.Equal(t, tt.expected, result, "CursorForward(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorBackward] function.
func Test_CursorBackward(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorBackward one", arg: 1, expected: "\x1b[D"},
		{name: "CursorBackward negative", arg: -1, expected: ""},
		{name: "CursorBackward zero", arg: 0, expected: ""},
		{name: "CursorBackward large", arg: 100, expected: "\x1b[100D"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorBackward(tt.arg)
			require.Equal(t, tt.expected, result, "CursorBackward(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorNextLineN] function.
func Test_CursorNextLineN(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorNextLineN", arg: 1, expected: "\x1b[E"},
		{name: "CursorNextLineN negative", arg: -1, expected: ""},
		{name: "CursorNextLineN zero", arg: 0, expected: ""},
		{name: "CursorNextLineN large", arg: 100, expected: "\x1b[100E"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorNextLineN(tt.arg)
			require.Equal(t, tt.expected, result, "CursorNextLineN(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorPreviousLineN] function.
func Test_CursorPreviousLineN(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorPreviousLineN", arg: 1, expected: "\x1b[F"},
		{name: "CursorPreviousLineN negative", arg: -1, expected: ""},
		{name: "CursorPreviousLineN zero", arg: 0, expected: ""},
		{name: "CursorPreviousLineN large", arg: 100, expected: "\x1b[100F"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorPreviousLineN(tt.arg)
			require.Equal(t, tt.expected, result, "CursorPreviousLineN(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorHorizontalAbsolute] function.
func Test_CursorHorizontalAbsolute(t *testing.T) {
	tests := []cursorTestCase{
		{name: "CursorHorizontalAbsolute one", arg: 1, expected: "\x1b[G"},
		{name: "CursorHorizontalAbsolute negative", arg: -1, expected: ""},
		{name: "CursorHorizontalAbsolute zero", arg: 0, expected: ""},
		{name: "CursorHorizontalAbsolute large", arg: 100, expected: "\x1b[100G"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorHorizontalAbsolute(tt.arg)
			require.Equal(t, tt.expected, result, "CursorHorizontalAbsolute(%d) = %q; want %q", tt.arg, result, tt.expected)
		})
	}
}

// Tests for [ansi.CursorPosition] function.
func Test_CursorPosition(t *testing.T) {
	tests := []struct {
		name        string
		row, column int
		expected    string
	}{
		{name: "CursorPosition 1,1", row: 1, column: 1, expected: "\x1b[H"},
		{name: "CursorPosition 10,20", row: 10, column: 20, expected: "\x1b[10;20H"},
		// Edge cases
		{name: "CursorPosition 0,0", row: 0, column: 0, expected: ""},                      // Both zero
		{name: "CursorPosition -1,-1", row: -1, column: -1, expected: ""},                  // Both negative
		{name: "CursorPosition 0,1", row: 0, column: 1, expected: ""},                      // Row zero
		{name: "CursorPosition 1,0", row: 1, column: 0, expected: ""},                      // Column zero
		{name: "CursorPosition -1,1", row: -1, column: 1, expected: ""},                    // Row negative
		{name: "CursorPosition 1,-1", row: 1, column: -1, expected: ""},                    // Column negative
		{name: "CursorPosition 100,200", row: 100, column: 200, expected: "\x1b[100;200H"}, // Large values
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ansi.CursorPosition(tt.row, tt.column)
			require.Equal(t, tt.expected, result, "CursorPosition(%d, %d) = %q; want %q", tt.row, tt.column, result, tt.expected)
		})
	}
}
