package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ansiCase struct {
	name        string
	input       string
	expectedLen int
}

var ansiCaseTests = []ansiCase{
	{name: "SGR reset", input: "\x1b[0m", expectedLen: 4},
	{name: "SGR red then text", input: "\x1b[31mhello", expectedLen: 5},
	{name: "CSI with valid intermediate", input: "\x1b[0 q", expectedLen: 5},
	{name: "OSC window title then BEL", input: "\x1b]0;My Title\x07", expectedLen: 13},
	{name: "OSC window title then ST", input: "\x1b]0;Title\x1b\\", expectedLen: 11},
	{name: "DCS with ST terminator", input: "\x1bPq#0;2;0;0;0\x1b\\", expectedLen: 15},
	{name: "DCS canceled by CAN", input: "\x1bPqdata\x18z", expectedLen: 7},
	{name: "SOS with ST terminator", input: "\x1bXhello\x1b\\", expectedLen: 9},
	{name: "PM with ST terminator", input: "\x1b^msg\x1b\\", expectedLen: 7},
	{name: "APC with ST terminator", input: "\x1b_data\x1b\\", expectedLen: 8},
	{name: "two-byte Fe", input: "\x1bD", expectedLen: 2},
	{name: "two-byte Fp", input: "\x1b7", expectedLen: 2},
	{name: "nF with multiple intermediates", input: "\x1b !Fx", expectedLen: 4},
	{name: "nF with invalid character", input: "\x1b 語!F", expectedLen: 0},
	{name: "malformed CSI remains split", input: "\x1b[ 1mok", expectedLen: 0},
	{name: "C1 CSI is not parsed", input: "\x9B31mhello", expectedLen: 0},
	{name: "7-bit OSC does not accept C1 ST", input: "\x1b]0;Title\x9Cz", expectedLen: 0},
	{name: "unterminated DCS", input: "\x1bPqpayload", expectedLen: 0},
	{name: "invalid escape sequence", input: "\x1b語", expectedLen: 0},
}

func Benchmark_EscapeLength(b *testing.B) {
	var tests = make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = EscapeLength(tt.input)
			}
		})
	}
}

// Tests for [EscapeLength] function.
func Test_EscapeLength(t *testing.T) {
	t.Parallel()

	errFormat := "EscapeLength(%q) returned incorrect length for type '%T'"
	var tests = make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, EscapeLength(tt.input), errFormat, tt.input, "")
			require.Equal(t, tt.expectedLen, EscapeLength(customStringType(tt.input)), errFormat, tt.input, customStringType(""))

			// []byte
			require.Equal(t, tt.expectedLen, EscapeLength([]byte(tt.input)), errFormat, tt.input, []byte(nil))
			require.Equal(t, tt.expectedLen, EscapeLength(customBytesType([]byte(tt.input))), errFormat, tt.input, customBytesType([]byte(nil)))

			// []rune
			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, EscapeLength(inputRunes), errFormat, tt.input, []rune(nil))
			require.Equal(t, tt.expectedLen, EscapeLength(customRunesType(inputRunes)), errFormat, tt.input, customRunesType([]rune(nil)))
		})
	}
}

// Tests for [escapeLength] function.
func Test_escapeLength(t *testing.T) {
	t.Parallel()

	errFormat := "escapeLength(%q) returned incorrect length for type '%T'"
	var tests = make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, escapeLength(tt.input), errFormat, tt.input, "")
			require.Equal(t, tt.expectedLen, escapeLength(customStringType(tt.input)), errFormat, tt.input, customStringType(""))

			// []byte
			require.Equal(t, tt.expectedLen, escapeLength([]byte(tt.input)), errFormat, tt.input, []byte(nil))
			require.Equal(t, tt.expectedLen, escapeLength(customBytesType([]byte(tt.input))), errFormat, tt.input, customBytesType([]byte(nil)))
		})
	}
}

// Tests for [csiBodyLength] function.
func Test_csiBodyLength(t *testing.T) {
	t.Parallel()

	errFormat := "csiBodyLength(%q) returned incorrect length for type '%T'"
	tests := []ansiCase{
		{name: "SGR reset", input: "\x1b[0m", expectedLen: 2},
		{name: "SGR red then text", input: "\x1b[31mhello", expectedLen: 3},
		{name: "CSI with valid intermediate", input: "\x1b[0 q", expectedLen: 3},
		{name: "CSI with multiple params", input: "\x1b[1;2;3m", expectedLen: 6},
		{name: "CSI with invalid character", input: "\x1b[語", expectedLen: 0},
		{name: "malformed CSI remains split", input: "\x1b[ 1mok", expectedLen: 0},
		{name: "empty CSI", input: "\x1b[", expectedLen: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, csiBodyLength(tt.input[2:]), errFormat, tt.input, "")
			require.Equal(t, tt.expectedLen, csiBodyLength(customStringType(tt.input[2:])), errFormat, tt.input, customStringType(""))

			// []byte
			require.Equal(t, tt.expectedLen, csiBodyLength([]byte(tt.input[2:])), errFormat, tt.input, []byte(nil))
			require.Equal(t, tt.expectedLen, csiBodyLength(customBytesType([]byte(tt.input[2:]))), errFormat, tt.input, customBytesType([]byte(nil)))
		})
	}
}

// Tests for [oscLength] function.
func Test_oscLength(t *testing.T) {
	t.Parallel()

	errFormat := "oscLength(%q) returned incorrect length"
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

			// string
			require.Equal(t, tt.expectedLen, oscLength(tt.input[2:]), errFormat, tt.input)
			require.Equal(t, tt.expectedLen, oscLength(customStringType(tt.input[2:])), errFormat, tt.input)

			// []byte
			require.Equal(t, tt.expectedLen, oscLength([]byte(tt.input[2:])), errFormat, tt.input)
			require.Equal(t, tt.expectedLen, oscLength(customBytesType([]byte(tt.input[2:]))), errFormat, tt.input)
		})
	}
}

// Tests for [stSequenceLength] function.
func Test_stSequenceLength(t *testing.T) {
	t.Parallel()

	errFormat := "stSequenceLength(%q) returned incorrect length"
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

			// string
			require.Equal(t, tt.expectedLen, stSequenceLength(tt.input[2:]), errFormat, tt.input)
			require.Equal(t, tt.expectedLen, stSequenceLength(customStringType(tt.input[2:])), errFormat, tt.input)

			// []byte
			require.Equal(t, tt.expectedLen, stSequenceLength([]byte(tt.input[2:])), errFormat, tt.input)
			require.Equal(t, tt.expectedLen, stSequenceLength(customBytesType([]byte(tt.input[2:]))), errFormat, tt.input)
		})
	}
}
