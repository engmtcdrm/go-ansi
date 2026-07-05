package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for [escapeLength8BitRune] function.
func Test_escapeLength8BitRune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "escapeLength8BitRune")
	tests := make([]ansiCase, len(ansi8BitCaseTests))
	copy(tests, ansi8BitCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input)
			require.Equal(t, tt.expectedLen, escapeLength8BitRune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, escapeLength8BitRune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [oscLengthC1Rune] function.
func Test_oscLengthC1Rune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "oscLengthC1Rune")
	tests := make([]ansiCase, len(osc8BitCaseTests))
	copy(tests, osc8BitCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input[1:])

			require.Equal(t, tt.expectedLen, oscLengthC1Rune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, oscLengthC1Rune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}

// Tests for [stSequenceLengthC1Rune] function.
func Test_stSequenceLengthC1Rune(t *testing.T) {
	t.Parallel()

	errFormat := getErrFormat(t, "stSequenceLengthC1Rune")
	tests := make([]ansiCase, len(stSequence8BitCaseTests))
	copy(tests, stSequence8BitCaseTests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputRunes := stringToRunes(t, tt.input[1:])
			require.Equal(t, tt.expectedLen, stSequenceLengthC1Rune(inputRunes), errFormat, tt.input, "[]rune")
			require.Equal(t, tt.expectedLen, stSequenceLengthC1Rune(customRunesType(inputRunes)), errFormat, tt.input, "customRunesType")
		})
	}
}
