package ansi

import (
	"testing"
)

// Tests for [escapeLengthRune] function.
func Test_escapeLengthRune(t *testing.T) {
	t.Parallel()

	var tests = make([]ansiCase, len(ansiCaseTests))
	copy(tests, ansiCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			returnedLen := escapeLengthRune([]rune(tt.input))
			if returnedLen != tt.expectedLen {
				t.Fatalf("escapeLengthRune returned %d, expected %d", returnedLen, tt.expectedLen)
			}
		})
	}
}

// Tests for [csiBodyLengthRune] function.
func Test_csiBodyLengthRune(t *testing.T) {
	t.Parallel()

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

			returnedLen := csiBodyLengthRune([]rune(tt.input[2:]))
			if returnedLen != tt.expectedLen {
				t.Fatalf("csiBodyLengthRune returned %d, expected %d", returnedLen, tt.expectedLen)
			}
		})
	}
}

// Tests for [oscLengthRune] function.
func Test_oscLengthRune(t *testing.T) {
	t.Parallel()

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

			returnedLen := oscLengthRune([]rune(tt.input[2:]))
			if returnedLen != tt.expectedLen {
				t.Fatalf("oscLengthRune returned %d, expected %d", returnedLen, tt.expectedLen)
			}
		})
	}
}

// Tests for [stSequenceLengthRune] function.
func Test_stSequenceLengthRune(t *testing.T) {
	t.Parallel()

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

			returnedLen := stSequenceLengthRune([]rune(tt.input[2:]))
			if returnedLen != tt.expectedLen {
				t.Fatalf("stSequenceLengthRune returned %d, expected %d", returnedLen, tt.expectedLen)
			}
		})
	}
}
