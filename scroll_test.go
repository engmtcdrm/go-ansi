package ansi_test

import (
	"testing"

	"github.com/engmtcdrm/go-ansi"
	"github.com/stretchr/testify/require"
)

// Tests for [ansi.ScrollUp] function.
func Test_ScrollUp(t *testing.T) {
	t.Run("ScrollUp", func(t *testing.T) {
		require.Equal(t, "\x1b[1S", ansi.ScrollUp(1), "ScrollUp(1) should return the correct ANSI escape code for scrolling up 1 line.")
	})

	t.Run("ScrollUp negative", func(t *testing.T) {
		require.Equal(t, "", ansi.ScrollUp(-1), "ScrollUp with a negative value should return an empty string.")
	})

	t.Run("ScrollUp zero", func(t *testing.T) {
		require.Equal(t, "", ansi.ScrollUp(0), "ScrollUp with a zero value should return an empty string.")
	})

	t.Run("ScrollUp large", func(t *testing.T) {
		require.Equal(t, "\x1b[100S", ansi.ScrollUp(100), "ScrollUp with a large value should return the correct ANSI escape code for scrolling up 100 lines.")
	})
}

// Tests for [ansi.ScrollDown] function.
func Test_ScrollDown(t *testing.T) {
	t.Run("ScrollDown", func(t *testing.T) {
		require.Equal(t, "\x1b[1T", ansi.ScrollDown(1), "ScrollDown(1) should return the correct ANSI escape code for scrolling down 1 line.")
	})

	t.Run("ScrollDown negative", func(t *testing.T) {
		require.Equal(t, "", ansi.ScrollDown(-1), "ScrollDown with a negative value should return an empty string.")
	})

	t.Run("ScrollDown zero", func(t *testing.T) {
		require.Equal(t, "", ansi.ScrollDown(0), "ScrollDown with a zero value should return an empty string.")
	})

	t.Run("ScrollDown large", func(t *testing.T) {
		require.Equal(t, "\x1b[100T", ansi.ScrollDown(100), "ScrollDown with a large value should return the correct ANSI escape code for scrolling down 999 lines.")
	})
}

// Tests for [ansi.ScrollUp1] and [ansi.ScrollDown1] to ensure they are correctly defined as derived constants.
func TestScrollDerivedConstants(t *testing.T) {
	require.Equal(t, ansi.ScrollUp(1), ansi.ScrollUp1, "ScrollUp1 should be equal to ScrollUp(1). ScrollUp")
	require.Equal(t, ansi.ScrollDown(1), ansi.ScrollDown1, "ScrollDown1 should be equal to ScrollDown(1). ScrollDown")
}
