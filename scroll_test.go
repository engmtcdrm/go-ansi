package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for [ScrollUp1] and [ScrollDown1] to ensure they are correctly defined as derived constants.
func Test_ScrollDerivedConstants(t *testing.T) {
	t.Parallel()

	require.Equal(t, ScrollUp(1), ScrollUp1, "ScrollUp1 should be equal to ScrollUp(1). ScrollUp")
	require.Equal(t, ScrollDown(1), ScrollDown1, "ScrollDown1 should be equal to ScrollDown(1). ScrollDown")
}

// Tests for [ScrollDown] function.
func Test_ScrollDown(t *testing.T) {
	t.Parallel()

	tests := []intToStringTestCase{
		{name: "-1", input: -1, expected: ""},
		{name: "0", input: 0, expected: ""},
		{name: "1", input: 1, expected: "\x1b[1T"},
		{name: "100", input: 100, expected: "\x1b[100T"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ScrollDown(tt.input)
			require.Equal(t, tt.expected, result, "ScrollDown(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}

// Tests for [ScrollUp] function.
func Test_ScrollUp(t *testing.T) {
	t.Parallel()

	tests := []intToStringTestCase{
		{name: "-1", input: -1, expected: ""},
		{name: "0", input: 0, expected: ""},
		{name: "1", input: 1, expected: "\x1b[1S"},
		{name: "100", input: 100, expected: "\x1b[100S"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ScrollUp(tt.input)
			require.Equal(t, tt.expected, result, "ScrollUp(%d) = %q; want %q", tt.input, result, tt.expected)
		})
	}
}
