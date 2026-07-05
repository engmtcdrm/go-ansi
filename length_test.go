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

var (
	// Test cases for [EscapeLength], [escapeLength], and [escapeLengthRune] functions.
	ansiCaseTests = []ansiCase{
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
		{name: "empty nF", input: "\x1b!", expectedLen: 0},
		{name: "nF with no intermediates", input: "\x1b!F", expectedLen: 3},
		{name: "nF with multiple intermediates", input: "\x1b !Fx", expectedLen: 4},
		{name: "nF with invalid character", input: "\x1b 語!F", expectedLen: 0},
		{name: "malformed CSI remains split", input: "\x1b[ 1mok", expectedLen: 0},
		{name: "C1 CSI is not parsed", input: "\x9B31mhello", expectedLen: 0},
		{name: "7-bit OSC does not accept C1 ST", input: "\x1b]0;Title\x9Cz", expectedLen: 0},
		{name: "unterminated DCS", input: "\x1bPqpayload", expectedLen: 0},
		{name: "invalid escape sequence", input: "\x1b語", expectedLen: 0},
	}

	// Test cases for [csiBodyLength] and [csiBodyLengthRune] functions.
	csiBodyCaseTests = []ansiCase{
		{name: "SGR reset", input: "\x1b[0m", expectedLen: 2},
		{name: "SGR red then text", input: "\x1b[31mhello", expectedLen: 3},
		{name: "CSI with valid intermediate", input: "\x1b[0 q", expectedLen: 3},
		{name: "CSI with multiple params", input: "\x1b[1;2;3m", expectedLen: 6},
		{name: "CSI with invalid character", input: "\x1b[語", expectedLen: 0},
		{name: "malformed CSI remains split", input: "\x1b[ 1mok", expectedLen: 0},
		{name: "empty CSI", input: "\x1b[", expectedLen: 0},
	}

	// Test cases for [intermediateThenFinalLength] and
	// [intermediateThenFinalLengthRune] functions.
	intermediateCaseTests = []ansiCase{
		{name: "empty nF", input: "\x1b!", expectedLen: 0},
		{name: "nF with no intermediates", input: "\x1b!F", expectedLen: 3},
		{name: "nF with multiple intermediates", input: "\x1b !Fx", expectedLen: 4},
		{name: "nF with invalid character", input: "\x1b 語!F", expectedLen: 0},
	}

	// Test cases for [oscLength] and [oscLengthRune] functions.
	oscCaseTests = []ansiCase{
		{name: "OSC window title then BEL", input: "\x1b]0;My Title\x07", expectedLen: 11},
		{name: "OSC window title then ST", input: "\x1b]0;Title\x1b\\", expectedLen: 9},
		{name: "OSC unterminated", input: "\x1b]0;Title", expectedLen: -1},
		{name: "OSC with cancel", input: "\x1b]0;My Title\x18", expectedLen: 10},
		{name: "OSC empty with cancel", input: "\x1b]\x18", expectedLen: 0},
	}

	// Test cases for [stSequenceLength] and [stSequenceLengthRune] functions.
	stSequenceCaseTests = []ansiCase{
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
)

// Benchmark for [EscapeLength] function.
func Benchmark_EscapeLength(b *testing.B) {
	tests := make([]ansiCase, len(ansiCaseTests))
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

	errFormat := getErrFormat(t, "EscapeLength")
	tests := make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, EscapeLength(tt.input), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, EscapeLength(customStringType(tt.input)), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input)
			require.Equal(t, tt.expectedLen, EscapeLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, EscapeLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")

			// []rune
			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, EscapeLength(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, EscapeLength(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [escapeLength] function.
func Test_escapeLength(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "escapeLength")
	tests := make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, escapeLength(tt.input), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, escapeLength(customStringType(tt.input)), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input)
			require.Equal(t, tt.expectedLen, escapeLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, escapeLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [csiBodyLength] function.
func Test_csiBodyLength(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "csiBodyLength")
	tests := make([]ansiCase, len(csiBodyCaseTests))
	copy(tests, csiBodyCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, csiBodyLength(tt.input[2:]), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, csiBodyLength(customStringType(tt.input[2:])), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input[2:])
			require.Equal(t, tt.expectedLen, csiBodyLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, csiBodyLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [intermediateThenFinalLength] function.
func Test_intermediateThenFinalLength(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "intermediateThenFinalLength")
	tests := make([]ansiCase, len(intermediateCaseTests))
	copy(tests, intermediateCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, intermediateThenFinalLength(tt.input), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, intermediateThenFinalLength(customStringType(tt.input)), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input)
			require.Equal(t, tt.expectedLen, intermediateThenFinalLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, intermediateThenFinalLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [oscLength] function.
func Test_oscLength(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "oscLength")
	tests := make([]ansiCase, len(oscCaseTests))
	copy(tests, oscCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, oscLength(tt.input[2:]), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, oscLength(customStringType(tt.input[2:])), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input[2:])
			require.Equal(t, tt.expectedLen, oscLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, oscLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}

// Tests for [stSequenceLength] function.
func Test_stSequenceLength(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "stSequenceLength")
	tests := make([]ansiCase, len(stSequenceCaseTests))
	copy(tests, stSequenceCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// string
			require.Equal(t, tt.expectedLen, stSequenceLength(tt.input[2:]), errFormat, tt.input, "string")
			require.Equal(t, tt.expectedLen, stSequenceLength(customStringType(tt.input[2:])), errFormat, tt.input, "customStringType")

			// []byte
			inputBytes := []byte(tt.input[2:])
			require.Equal(t, tt.expectedLen, stSequenceLength(inputBytes), errFormat, tt.input, "[]byte")
			require.Equal(t, tt.expectedLen, stSequenceLength(customBytesType(inputBytes)), errFormat, tt.input, "customBytesType")
		})
	}
}
