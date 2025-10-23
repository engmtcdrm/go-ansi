package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestScroll(t *testing.T) {
	tests := []struct {
		name     string
		function func(int) string
		arg      int
		expected string
	}{
		{"ScrollUpN", ansi.ScrollUpN, 1, "\x1b[1S"},
		{"ScrollDownN", ansi.ScrollDownN, 1, "\x1b[1T"},
		// Edge cases - negative values
		{"ScrollUpN negative", ansi.ScrollUpN, -1, ""},
		{"ScrollDownN negative", ansi.ScrollDownN, -1, ""},
		// Edge cases - zero values
		{"ScrollUpN zero", ansi.ScrollUpN, 0, ""},
		{"ScrollDownN zero", ansi.ScrollDownN, 0, ""},
		// Large values
		{"ScrollUpN large", ansi.ScrollUpN, 100, "\x1b[100S"},
		{"ScrollDownN large", ansi.ScrollDownN, 999, "\x1b[999T"},
	}

	for _, test := range tests {
		result := test.function(test.arg)
		if result != test.expected {
			t.Errorf("%s(%d) = %q; want %q", test.name, test.arg, result, test.expected)
		}
	}
}

func TestScrollDerivedConstants(t *testing.T) {
	// Test composite/derived constants
	if ansi.ScrollUp1 != ansi.ScrollUpN(1) {
		t.Errorf("ScrollUp1 = %q; want %q", ansi.ScrollUp1, ansi.ScrollUpN(1))
	}
	if ansi.ScrollDown1 != ansi.ScrollDownN(1) {
		t.Errorf("ScrollDown1 = %q; want %q", ansi.ScrollDown1, ansi.ScrollDownN(1))
	}
}
