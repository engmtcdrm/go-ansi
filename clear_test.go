package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
)

func TestClearLineReset(t *testing.T) {
	// Test composite constant: ClearLineReset should be ClearLine + CursorLineBegin
	expected := ansi.ClearLine + ansi.CursorLineBegin
	if ansi.ClearLineReset != expected {
		t.Errorf("ClearLineReset = %q; want %q (ClearLine + CursorLineBegin)", ansi.ClearLineReset, expected)
	}
}
