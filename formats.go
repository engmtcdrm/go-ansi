package ansi

// Formatting

const (
	Reset           = CSI + "0m"  // Reset all formatting.
	Bold            = CSI + "1m"  // Bold or increased intensity.
	Dim             = CSI + "2m"  // Dim, decreased intensity, or dim.
	Italic          = CSI + "3m"  // Italicized text.
	Underline       = CSI + "4m"  // Underlined text.
	SlowBlink       = CSI + "5m"  // Slow blinking text.
	RapidBlink      = CSI + "6m"  // Rapid blinking text.
	Reverse         = CSI + "7m"  // Reversed colors. Not widely supported.
	Hide            = CSI + "8m"  // Hide text. Not widely supported.
	Strike          = CSI + "9m"  // Strikethrough text.
	DoubleUnderline = CSI + "21m" // Double underlined text.

	// Reset Formatting

	ResetIntensity = CSI + "22m" // Reset bold or faint.
	ResetItalic    = CSI + "23m" // Reset italic.
	ResetUnderline = CSI + "24m" // Reset underline.
	ResetBlink     = CSI + "25m" // Reset blink.
	ResetReverse   = CSI + "27m" // Reset reverse.
	ResetHide      = CSI + "28m" // Reset hide.
	ResetStrike    = CSI + "29m" // Reset strikethrough.

	Framed         = CSI + "51m" // Framed text.
	Encircled      = CSI + "52m" // Encircled text.
	Overlined      = CSI + "53m" // Overlined text.
	ResetFramed    = CSI + "54m" // Reset framed and encircled.
	ResetOverlined = CSI + "55m" // Reset overlined.
	IdeogramRight  = CSI + "73m" // Ideogram right underline.

	// Hidden text. Not widely supported.
	//
	// Deprecated: use [Hide] instead. This will be removed in v2.0.0.
	Hidden = Hide

	// Faint, decreased intensity, or dim.
	//
	// Deprecated: use [Dim] instead. This will be removed in v2.0.0.
	Faint = Dim

	// Reset Hidden text. Not widely supported.
	//
	// Deprecated: use [ResetHide] instead. This will be removed in v2.0.0.
	ResetHidden = ResetHide
)
