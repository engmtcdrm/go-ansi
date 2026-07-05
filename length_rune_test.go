package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for [escapeLengthRune] function.
func Test_escapeLengthRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "escapeLengthRune")
	var tests = make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, escapeLengthRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, escapeLengthRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [csiBodyLengthRune] function.
func Test_csiBodyLengthRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "csiBodyLengthRune")
	var tests = make([]ansiCase, len(csiBodyCaseTests))
	copy(tests, csiBodyCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input[2:])
			require.Equal(t, tt.expectedLen, csiBodyLengthRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, csiBodyLengthRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [intermediateThenFinalLengthRune] function.
func Test_intermediateThenFinalLengthRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "intermediateThenFinalLengthRune")
	var tests = make([]ansiCase, len(intermediateCaseTests))
	copy(tests, intermediateCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, intermediateThenFinalLengthRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, intermediateThenFinalLengthRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [oscLengthRune] function.
func Test_oscLengthRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "oscLengthRune")
	tests := []ansiCase{
		{name: "OSC window title then BEL", input: "\x1b]0;My Title\x07", expectedLen: 11},
		{name: "OSC window title then ST", input: "\x1b]0;Title\x1b\\", expectedLen: 9},
		{name: "OSC unterminated", input: "\x1b]0;Title", expectedLen: -1},
		{name: "OSC with cancel", input: "\x1b]0;My Title\x18", expectedLen: 10},
		{name: "OSC empty with cancel", input: "\x1b]\x18", expectedLen: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input[2:])
			require.Equal(t, tt.expectedLen, oscLengthRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, oscLengthRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [stSequenceLengthRune] function.
func Test_stSequenceLengthRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "stSequenceLengthRune")
	tests := []ansiCase{
		{name: "DCS with ST terminator", input: "\x1bPq#0;2;0;0;0\x1b\\", expectedLen: 13},
		{name: "DCS canceled by CAN", input: "\x1bPqdata\x18z", expectedLen: 5},
		{name: "SOS with ST terminator", input: "\x1bXhello\x1b\\", expectedLen: 7},
		{name: "PM with ST terminator", input: "\x1b^msg\x1b\\", expectedLen: 5},
		{name: "APC with ST terminator", input: "\x1b_data\x1b\\", expectedLen: 6},
		{name: "unterminated DCS", input: "\x1bPqpayload", expectedLen: -1},
		{name: "unterminated SOS", input: "\x1bXhello", expectedLen: -1},
		{name: "unterminated PM", input: "\x1b^msg", expectedLen: -1},
		{name: "unterminated APC", input: "\x1b_data", expectedLen: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input[2:])
			require.Equal(t, tt.expectedLen, stSequenceLengthRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, stSequenceLengthRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// stringToRunes is a helper function that converts a string into a slice of
// runes. You cannot always do a directly conversion from string to []rune
// because ANSI 8-bit escape sequences may contain bytes that do not correspond
// to valid UTF-8 runes, and directly converting a string to []rune could
// misinterpret these bytes.
func stringToRunes(t *testing.T, s string) []rune {
	t.Helper()

	var runes []rune
	for _, b := range []byte(s) {
		runes = append(runes, rune(b))
	}

	return runes
}
