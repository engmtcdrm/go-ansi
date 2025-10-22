package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestFaintDeprecated(t *testing.T) {
	// Test that deprecated Faint equals Dim
	if ansi.Faint != ansi.Dim {
		t.Errorf("Faint = %q; want %q (same as Dim)", ansi.Faint, ansi.Dim)
	}
}
