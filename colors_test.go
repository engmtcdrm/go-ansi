package ansi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type color8BitTestCase struct {
	name     string
	color    int
	expected string
}

type color24BitTestCase struct {
	name     string
	r, g, b  int
	expected string
}

// Tests for [Foreground8Bit] function.
func Test_Foreground8Bit(t *testing.T) {
	t.Parallel()

	tests := []color8BitTestCase{
		{name: "-1", color: -1, expected: ""},
		{name: "0", color: 0, expected: "\x1b[38;5;0m"},
		{name: "1", color: 1, expected: "\x1b[38;5;1m"},
		{name: "128", color: 128, expected: "\x1b[38;5;128m"},
		{name: "254", color: 254, expected: "\x1b[38;5;254m"},
		{name: "255", color: 255, expected: "\x1b[38;5;255m"},
		{name: "256", color: 256, expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Foreground8Bit(tt.color)
			require.Equal(t, tt.expected, result, "Foreground8Bit(%d) = %q; want %q", tt.color, result, tt.expected)
		})
	}
}

// Tests for [Background8Bit] function.
func Test_Background8Bit(t *testing.T) {
	t.Parallel()

	tests := []color8BitTestCase{
		{name: "-1", color: -1, expected: ""},
		{name: "0", color: 0, expected: "\x1b[48;5;0m"},
		{name: "1", color: 1, expected: "\x1b[48;5;1m"},
		{name: "128", color: 128, expected: "\x1b[48;5;128m"},
		{name: "254", color: 254, expected: "\x1b[48;5;254m"},
		{name: "255", color: 255, expected: "\x1b[48;5;255m"},
		{name: "256", color: 256, expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Background8Bit(tt.color)
			require.Equal(t, tt.expected, result, "Background8Bit(%d) = %q; want %q", tt.color, result, tt.expected)
		})
	}
}

// Tests for [Foreground24Bit] function.
func Test_Foreground24Bit(t *testing.T) {
	t.Parallel()

	tests := []color24BitTestCase{
		{name: "-1,0,0", r: -1, g: 0, b: 0, expected: ""},
		{name: "0,-1,0", r: 0, g: -1, b: 0, expected: ""},
		{name: "0,0,-1", r: 0, g: 0, b: -1, expected: ""},
		{name: "-1,-1,-1", r: -1, g: -1, b: -1, expected: ""},
		{name: "0,0,0 - Black", r: 0, g: 0, b: 0, expected: "\x1b[38;2;0;0;0m"},
		{name: "1,1,1", r: 1, g: 1, b: 1, expected: "\x1b[38;2;1;1;1m"},
		{name: "128,64,32", r: 128, g: 64, b: 32, expected: "\x1b[38;2;128;64;32m"},
		{name: "254,254,254", r: 254, g: 254, b: 254, expected: "\x1b[38;2;254;254;254m"},
		{name: "255,0,0 - Red", r: 255, g: 0, b: 0, expected: "\x1b[38;2;255;0;0m"},
		{name: "0,255,0 - Green", r: 0, g: 255, b: 0, expected: "\x1b[38;2;0;255;0m"},
		{name: "0,0,255 - Blue", r: 0, g: 0, b: 255, expected: "\x1b[38;2;0;0;255m"},
		{name: "255,128,0 - Orange", r: 255, g: 128, b: 0, expected: "\x1b[38;2;255;128;0m"},
		{name: "255,255,255 - White", r: 255, g: 255, b: 255, expected: "\x1b[38;2;255;255;255m"},
		{name: "256,0,0", r: 256, g: 0, b: 0, expected: ""},
		{name: "0,256,0", r: 0, g: 256, b: 0, expected: ""},
		{name: "0,0,256", r: 0, g: 0, b: 256, expected: ""},
		{name: "256,256,256", r: 256, g: 256, b: 256, expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Foreground24Bit(tt.r, tt.g, tt.b)
			require.Equal(t, tt.expected, result, "Foreground24Bit(%d, %d, %d) = %q; want %q", tt.r, tt.g, tt.b, result, tt.expected)
		})
	}
}

// Tests for [Background24Bit] function.
func Test_Background24Bit(t *testing.T) {
	t.Parallel()

	tests := []color24BitTestCase{
		{name: "-1,0,0", r: -1, g: 0, b: 0, expected: ""},
		{name: "0,-1,0", r: 0, g: -1, b: 0, expected: ""},
		{name: "0,0,-1", r: 0, g: 0, b: -1, expected: ""},
		{name: "-1,-1,-1", r: -1, g: -1, b: -1, expected: ""},
		{name: "0,0,0 - Black", r: 0, g: 0, b: 0, expected: "\x1b[48;2;0;0;0m"},
		{name: "1,1,1", r: 1, g: 1, b: 1, expected: "\x1b[48;2;1;1;1m"},
		{name: "128,64,32", r: 128, g: 64, b: 32, expected: "\x1b[48;2;128;64;32m"},
		{name: "254,254,254", r: 254, g: 254, b: 254, expected: "\x1b[48;2;254;254;254m"},
		{name: "255,0,0 - Red", r: 255, g: 0, b: 0, expected: "\x1b[48;2;255;0;0m"},
		{name: "0,255,0 - Green", r: 0, g: 255, b: 0, expected: "\x1b[48;2;0;255;0m"},
		{name: "0,0,255 - Blue", r: 0, g: 0, b: 255, expected: "\x1b[48;2;0;0;255m"},
		{name: "255,128,0 - Orange", r: 255, g: 128, b: 0, expected: "\x1b[48;2;255;128;0m"},
		{name: "255,255,255 - White", r: 255, g: 255, b: 255, expected: "\x1b[48;2;255;255;255m"},
		{name: "256,0,0", r: 256, g: 0, b: 0, expected: ""},
		{name: "0,256,0", r: 0, g: 256, b: 0, expected: ""},
		{name: "0,0,256", r: 0, g: 0, b: 256, expected: ""},
		{name: "256,256,256", r: 256, g: 256, b: 256, expected: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := Background24Bit(test.r, test.g, test.b)
			require.Equal(t, test.expected, result, "Background24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		})
	}
}
