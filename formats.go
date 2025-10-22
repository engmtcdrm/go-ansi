package ansi

// Formatting

const (
	Reset           = csi + "0m"  // Reset all formatting.
	Bold            = csi + "1m"  // Bold or increased intensity.
	Dim             = csi + "2m"  // Dim, decreased intensity, or dim.
	Italic          = csi + "3m"  // Italicized text.
	Underline       = csi + "4m"  // Underlined text.
	SlowBlink       = csi + "5m"  // Slow blinking text.
	RapidBlink      = csi + "6m"  // Rapid blinking text.
	Reverse         = csi + "7m"  // Reversed colors.
	Hidden          = csi + "8m"  // Hidden text.
	Strike          = csi + "9m"  // Strikethrough text.
	DoubleUnderline = csi + "21m" // Double underlined text.

	// Faint, decreased intensity, or dim.
	//
	// Deprecated: use [Dim] instead. This will be removed in v2.0.0.
	Faint = Dim

	// Reset Formatting

	ResetIntensity = csi + "22m" // Reset bold or faint.
	ResetItalic    = csi + "23m" // Reset italic.
	ResetUnderline = csi + "24m" // Reset underline.
	ResetBlink     = csi + "25m" // Reset blink.
	ResetReverse   = csi + "27m" // Reset reverse.
	ResetHidden    = csi + "28m" // Reset hidden.
	ResetStrike    = csi + "29m" // Reset strikethrough.
)
