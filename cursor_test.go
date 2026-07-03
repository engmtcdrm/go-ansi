package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type intToStringTestCase struct {
	name     string
	input    int
	expected string
}

// Tests for [CursorTopLeft] variable.
func Test_CursorTopLeft(t *testing.T) {
	t.Parallel()

	expected := CursorPosition(1, 1)
	require.Equal(t, expected, CursorTopLeft, "CursorTopLeft should be equal to its expected value.")
}

// Tests for [CursorLineBegin] variable.
func Test_CursorLineBegin(t *testing.T) {
	t.Parallel()

	expected := CursorHorizontalAbsolute(1)
	require.Equal(t, expected, CursorLineBegin, "CursorLineBegin should be equal to its expected value.")
}

// Tests for [CursorNextLine] variable.
func Test_CursorNextLine(t *testing.T) {
	t.Parallel()

	expected := CursorNextLineN(1)
	require.Equal(t, expected, CursorNextLine, "CursorNextLine should be equal to its expected value.")
}

// Tests for [CursorPreviousLine] variable.
func Test_CursorPreviousLine(t *testing.T) {
	t.Parallel()

	expected := CursorPreviousLineN(1)
	require.Equal(t, expected, CursorPreviousLine, "CursorPreviousLine should be equal to its expected value.")
}

// Tests for [CursorUp] function.
func Test_CursorUp(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorUp one", input: 1, expected: "\x1b[A"},
		{name: "CursorUp negative", input: -1, expected: ""},
		{name: "CursorUp zero", input: 0, expected: ""},
		{name: "CursorUp large", input: 100, expected: "\x1b[100A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorUp(tt.input)
			require.Equal(t, tt.expected, result, "CursorUp(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorDown] function.
func Test_CursorDown(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorDown one", input: 1, expected: "\x1b[B"},
		{name: "CursorDown negative", input: -1, expected: ""},
		{name: "CursorDown zero", input: 0, expected: ""},
		{name: "CursorDown large", input: 100, expected: "\x1b[100B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorDown(tt.input)
			require.Equal(t, tt.expected, result, "CursorDown(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorForward] function.
func Test_CursorForward(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorForward one", input: 1, expected: "\x1b[C"},
		{name: "CursorForward negative", input: -1, expected: ""},
		{name: "CursorForward zero", input: 0, expected: ""},
		{name: "CursorForward large", input: 100, expected: "\x1b[100C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorForward(tt.input)
			require.Equal(t, tt.expected, result, "CursorForward(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorBackward] function.
func Test_CursorBackward(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorBackward one", input: 1, expected: "\x1b[D"},
		{name: "CursorBackward negative", input: -1, expected: ""},
		{name: "CursorBackward zero", input: 0, expected: ""},
		{name: "CursorBackward large", input: 100, expected: "\x1b[100D"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorBackward(tt.input)
			require.Equal(t, tt.expected, result, "CursorBackward(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorNextLineN] function.
func Test_CursorNextLineN(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorNextLineN", input: 1, expected: "\x1b[E"},
		{name: "CursorNextLineN negative", input: -1, expected: ""},
		{name: "CursorNextLineN zero", input: 0, expected: ""},
		{name: "CursorNextLineN large", input: 100, expected: "\x1b[100E"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorNextLineN(tt.input)
			require.Equal(t, tt.expected, result, "CursorNextLineN(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorPreviousLineN] function.
func Test_CursorPreviousLineN(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorPreviousLineN", input: 1, expected: "\x1b[F"},
		{name: "CursorPreviousLineN negative", input: -1, expected: ""},
		{name: "CursorPreviousLineN zero", input: 0, expected: ""},
		{name: "CursorPreviousLineN large", input: 100, expected: "\x1b[100F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorPreviousLineN(tt.input)
			require.Equal(t, tt.expected, result, "CursorPreviousLineN(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorHorizontalAbsolute] function.
func Test_CursorHorizontalAbsolute(t *testing.T) {
	tests := []intToStringTestCase{
		{name: "CursorHorizontalAbsolute one", input: 1, expected: "\x1b[G"},
		{name: "CursorHorizontalAbsolute negative", input: -1, expected: ""},
		{name: "CursorHorizontalAbsolute zero", input: 0, expected: ""},
		{name: "CursorHorizontalAbsolute large", input: 100, expected: "\x1b[100G"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorHorizontalAbsolute(tt.input)
			require.Equal(t, tt.expected, result, "CursorHorizontalAbsolute(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [CursorPosition] function.
func Test_CursorPosition(t *testing.T) {
	tests := []struct {
		name        string
		row, column int
		expected    string
	}{
		{name: "CursorPosition 1,1", row: 1, column: 1, expected: "\x1b[H"},
		{name: "CursorPosition 10,20", row: 10, column: 20, expected: "\x1b[10;20H"},
		{name: "CursorPosition 0,0", row: 0, column: 0, expected: ""},                      // Both zero
		{name: "CursorPosition -1,-1", row: -1, column: -1, expected: ""},                  // Both negative
		{name: "CursorPosition 0,1", row: 0, column: 1, expected: ""},                      // Row zero
		{name: "CursorPosition 1,0", row: 1, column: 0, expected: ""},                      // Column zero
		{name: "CursorPosition -1,1", row: -1, column: 1, expected: ""},                    // Row negative
		{name: "CursorPosition 1,-1", row: 1, column: -1, expected: ""},                    // Column negative
		{name: "CursorPosition 100,200", row: 100, column: 200, expected: "\x1b[100;200H"}, // Large values
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CursorPosition(tt.row, tt.column)
			require.Equal(t, tt.expected, result, "CursorPosition(%d, %d) = %q; want %q", tt.row, tt.column, result, tt.expected)
		})
	}
}
