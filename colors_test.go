package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestForeground8Bit(t *testing.T) {
	tests := []struct {
		color    int
		expected string
	}{
		{0, "\x1b[38;5;0m"},
		{255, "\x1b[38;5;255m"},
		{-1, ""},
		{256, ""},
		// Additional edge cases
		{128, "\x1b[38;5;128m"}, // Mid-range value
		{1, "\x1b[38;5;1m"},     // Minimum valid positive
		{254, "\x1b[38;5;254m"}, // Maximum - 1
		{-100, ""},              // Large negative
		{1000, ""},              // Large positive
	}

	for _, test := range tests {
		result := ansi.Foreground8Bit(test.color)
		if result != test.expected {
			t.Errorf("Foreground8Bit(%d) = %q; want %q", test.color, result, test.expected)
		}
	}
}

func TestBackground8Bit(t *testing.T) {
	tests := []struct {
		color    int
		expected string
	}{
		{0, "\x1b[48;5;0m"},
		{255, "\x1b[48;5;255m"},
		{-1, ""},
		{256, ""},
		// Additional edge cases
		{128, "\x1b[48;5;128m"}, // Mid-range value
		{1, "\x1b[48;5;1m"},     // Minimum valid positive
		{254, "\x1b[48;5;254m"}, // Maximum - 1
		{-100, ""},              // Large negative
		{1000, ""},              // Large positive
	}

	for _, test := range tests {
		result := ansi.Background8Bit(test.color)
		if result != test.expected {
			t.Errorf("Background8Bit(%d) = %q; want %q", test.color, result, test.expected)
		}
	}
}

func TestForeground24Bit(t *testing.T) {
	tests := []struct {
		r, g, b  int
		expected string
	}{
		{0, 0, 0, "\x1b[38;2;0;0;0m"},
		{255, 255, 255, "\x1b[38;2;255;255;255m"},
		{-1, 0, 0, ""},
		{0, -1, 0, ""},
		{0, 0, -1, ""},
		{256, 0, 0, ""},
		{0, 256, 0, ""},
		{0, 0, 256, ""},
		// Additional edge cases
		{128, 64, 32, "\x1b[38;2;128;64;32m"},     // Mid-range values
		{255, 0, 0, "\x1b[38;2;255;0;0m"},         // Pure red
		{0, 255, 0, "\x1b[38;2;0;255;0m"},         // Pure green
		{0, 0, 255, "\x1b[38;2;0;0;255m"},         // Pure blue
		{1, 1, 1, "\x1b[38;2;1;1;1m"},             // Minimum valid positive
		{254, 254, 254, "\x1b[38;2;254;254;254m"}, // Maximum - 1
		{-1, -1, -1, ""},                          // All negative
		{256, 256, 256, ""},                       // All over max
		{255, 128, 0, "\x1b[38;2;255;128;0m"},     // Orange
	}

	for _, test := range tests {
		result := ansi.Foreground24Bit(test.r, test.g, test.b)
		if result != test.expected {
			t.Errorf("Foreground24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		}
	}
}

func TestBackground24Bit(t *testing.T) {
	tests := []struct {
		r, g, b  int
		expected string
	}{
		{0, 0, 0, "\x1b[48;2;0;0;0m"},
		{255, 255, 255, "\x1b[48;2;255;255;255m"},
		{-1, 0, 0, ""},
		{0, -1, 0, ""},
		{0, 0, -1, ""},
		{256, 0, 0, ""},
		{0, 256, 0, ""},
		{0, 0, 256, ""},
		// Additional edge cases
		{128, 64, 32, "\x1b[48;2;128;64;32m"},     // Mid-range values
		{255, 0, 0, "\x1b[48;2;255;0;0m"},         // Pure red
		{0, 255, 0, "\x1b[48;2;0;255;0m"},         // Pure green
		{0, 0, 255, "\x1b[48;2;0;0;255m"},         // Pure blue
		{1, 1, 1, "\x1b[48;2;1;1;1m"},             // Minimum valid positive
		{254, 254, 254, "\x1b[48;2;254;254;254m"}, // Maximum - 1
		{-1, -1, -1, ""},                          // All negative
		{256, 256, 256, ""},                       // All over max
		{255, 128, 0, "\x1b[48;2;255;128;0m"},     // Orange
	}

	for _, test := range tests {
		result := ansi.Background24Bit(test.r, test.g, test.b)
		if result != test.expected {
			t.Errorf("Background24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		}
	}
}
