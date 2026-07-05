package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var ansiCase8BitTests = []ansiCase{
	{name: "empty input", input: "", expectedLen: 0},
	{name: "C1 CSI with empty body", input: "\x9B", expectedLen: 0},
	{name: "C1 CSI then text", input: "\x9B31mhello", expectedLen: 4},
	{name: "C1 CSI multiple params", input: "\x9B1;2;3m", expectedLen: 7},
	{name: "C1 OSC with C1 ST", input: "\x9D0;Title\x9C", expectedLen: 9},
	{name: "C1 OSC with 7-bit ST is not parsed as one sequence", input: "\x9D0;Title\x1b\\", expectedLen: 0},
	{name: "C1 DCS with C1 ST", input: "\x90qpayload\x9C", expectedLen: 10},
	{name: "C1 DCS with 7-bit ST is not parsed as one sequence", input: "\x90qpayload\x1b\\", expectedLen: 0},
	{name: "C1 DCS canceled by CAN", input: "\x90qpayload\x18x", expectedLen: 9},
	{name: "C1 SOS with C1 ST", input: "\x98hello\x9C", expectedLen: 7},
	{name: "C1 PM with 7-bit ST is not parsed as one sequence", input: "\x9Emsg\x1b\\", expectedLen: 0},
	{name: "C1 APC with C1 ST", input: "\x9Fdata\x9C", expectedLen: 6},
	{name: "single C1 Fe control", input: "\x84", expectedLen: 1},
	{name: "C1 OSC unterminated", input: "\x9D0;title", expectedLen: 0},
	{name: "C1 DCS unterminated", input: "\x90data", expectedLen: 0},
	{name: "7-bit ESC sequence is not parsed", input: "\x1b[31mhello", expectedLen: 0},
}

// Tests for [EscapeLength8Bit] function.
func Test_EscapeLength8Bit(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "EscapeLength8Bit")
	var tests = make([]ansiCase, len(ansiCase8BitTests))
	copy(tests, ansiCase8BitTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// strings
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(tt.input), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(customStringType(tt.input)), errFormat, tt.input, "customStringType")

			// []bytes
			inputBytes := []byte(tt.input)
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")

			// []runes
			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, EscapeLength8Bit(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [escapeLength8Bit] function.
func Test_escapeLength8Bit(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "escapeLength8Bit")
	var tests = make([]ansiCase, len(ansiCase8BitTests))
	copy(tests, ansiCase8BitTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// strings
			require.Equal(t, tt.expectedLen, escapeLength8Bit(tt.input), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, escapeLength8Bit(customStringType(tt.input)), errFormat, tt.input, "customStringType")

			// []bytes
			inputBytes := []byte(tt.input)
			require.Equal(t, tt.expectedLen, escapeLength8Bit(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, escapeLength8Bit(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [oscLengthC1] function.
func Test_oscLengthC1(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "oscLengthC1")
	tests := []ansiCase{
		{name: "OSC empty input", input: "\x9D", expectedLen: -1},
		{name: "OSC with BEL terminator", input: "\x9D0;Title\x07", expectedLen: 8},
		{name: "OSC with C1 ST terminator", input: "\x9D0;Title\x9C", expectedLen: 8},
		{name: "OSC with cancel", input: "\x9D0;Title\x18", expectedLen: 7},
		{name: "OSC unterminated", input: "\x9D0;Title", expectedLen: -1},
		{name: "OSC empty with cancel", input: "\x9D\x18", expectedLen: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// strings
			require.Equal(t, tt.expectedLen, oscLengthC1(tt.input[1:]), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, oscLengthC1(customStringType(tt.input[1:])), errFormat, tt.input, "customStringType")

			// []bytes
			inputBytes := []byte(tt.input[1:])
			require.Equal(t, tt.expectedLen, oscLengthC1(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, oscLengthC1(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [stSequenceLengthC1] function.
func Test_stSequenceLengthC1(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "stSequenceLengthC1")
	tests := []ansiCase{
		{name: "DCS with C1 ST terminator", input: "\x90qpayload\x9C", expectedLen: 9},
		{name: "DCS canceled by CAN", input: "\x90qpayload\x18x", expectedLen: 8},
		{name: "SOS with C1 ST terminator", input: "\x98hello\x9C", expectedLen: 6},
		{name: "PM with C1 ST terminator", input: "\x9Emsg\x9C", expectedLen: 4},
		{name: "APC with C1 ST terminator", input: "\x9Fdata\x9C", expectedLen: 5},
		{name: "unterminated DCS", input: "\x90qpayload", expectedLen: -1},
		{name: "unterminated SOS", input: "\x98hello", expectedLen: -1},
		{name: "unterminated PM", input: "\x9Emsg", expectedLen: -1},
		{name: "unterminated APC", input: "\x9Fdata", expectedLen: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// strings
			require.Equal(t, tt.expectedLen, stSequenceLengthC1(tt.input[1:]), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, stSequenceLengthC1(customStringType(tt.input[1:])), errFormat, tt.input, "customStringType")

			// []bytes
			inputBytes := []byte(tt.input[1:])
			require.Equal(t, tt.expectedLen, stSequenceLengthC1(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, stSequenceLengthC1(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}
