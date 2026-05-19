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
	Reverse         = csi + "7m"  // Reversed colors. Not widely supported.
	Hide            = csi + "8m"  // Hide text. Not widely supported.
	Strike          = csi + "9m"  // Strikethrough text.
	DoubleUnderline = csi + "21m" // Double underlined text.

	// Reset Formatting

	ResetIntensity = csi + "22m" // Reset bold or faint.
	ResetItalic    = csi + "23m" // Reset italic.
	ResetUnderline = csi + "24m" // Reset underline.
	ResetBlink     = csi + "25m" // Reset blink.
	ResetReverse   = csi + "27m" // Reset reverse.
	ResetHide      = csi + "28m" // Reset hide.
	ResetStrike    = csi + "29m" // Reset strikethrough.

	Framed         = csi + "51m" // Framed text.
	Encircled      = csi + "52m" // Encircled text.
	Overlined      = csi + "53m" // Overlined text.
	ResetFramed    = csi + "54m" // Reset framed and encircled.
	ResetOverlined = csi + "55m" // Reset overlined.
	IdeogramRight  = csi + "73m" // Ideogram right underline.

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
