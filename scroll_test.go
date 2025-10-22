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
	}

	for _, test := range tests {
		result := test.function(test.arg)
		if result != test.expected {
			t.Errorf("%s(%d) = %q; want %q", test.name, test.arg, result, test.expected)
		}
	}
}
