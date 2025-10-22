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
	}

	for _, test := range tests {
		result := ansi.Background24Bit(test.r, test.g, test.b)
		if result != test.expected {
			t.Errorf("Background24Bit(%d, %d, %d) = %q; want %q", test.r, test.g, test.b, result, test.expected)
		}
	}
}
