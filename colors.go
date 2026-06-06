package ansi

import "strconv"

// ANSI color codes.
const (
	ANSIBlack = iota
	ANSIRed
	ANSIGreen
	ANSIYellow
	ANSIBlue
	ANSIMagenta
	ANSICyan
	ANSIWhite
	ANSIBrightBlack
	ANSIBrightRed
	ANSIBrightGreen
	ANSIBrightYellow
	ANSIBrightBlue
	ANSIBrightMagenta
	ANSIBrightCyan
	ANSIBrightWhite
)

// Foreground Color

const (
	Black   = CSI + "30m" // Black foreground color.
	Red     = CSI + "31m" // Red foreground color.
	Green   = CSI + "32m" // Green foreground color.
	Yellow  = CSI + "33m" // Yellow foreground color.
	Blue    = CSI + "34m" // Blue foreground color.
	Magenta = CSI + "35m" // Magenta foreground color.
	Cyan    = CSI + "36m" // Cyan foreground color.
	White   = CSI + "37m" // White foreground color.
	Default = CSI + "39m" // Default foreground color.
)

// Background Color

const (
	BlackBg   = CSI + "40m" // Black background color.
	RedBg     = CSI + "41m" // Red background color.
	GreenBg   = CSI + "42m" // Green background color.
	YellowBg  = CSI + "43m" // Yellow background color.
	BlueBg    = CSI + "44m" // Blue background color.
	MagentaBg = CSI + "45m" // Magenta background color.
	CyanBg    = CSI + "46m" // Cyan background color.
	WhiteBg   = CSI + "47m" // White background color.
	DefaultBg = CSI + "49m" // Default background color.
)

// Intense Foreground Color

const (
	IntenseBlack   = CSI + "90m" // Intense Black foreground color.
	IntenseRed     = CSI + "91m" // Intense Red foreground color.
	IntenseGreen   = CSI + "92m" // Intense Green foreground color.
	IntenseYellow  = CSI + "93m" // Intense Yellow foreground color.
	IntenseBlue    = CSI + "94m" // Intense Blue foreground color.
	IntenseMagenta = CSI + "95m" // Intense Magenta foreground color.
	IntenseCyan    = CSI + "96m" // Intense Cyan foreground color.
	IntenseWhite   = CSI + "97m" // Intense White foreground color.

)

// Intense Background Color

const (
	IntenseBlackBg   = CSI + "100m" // Intense Black background color.
	IntenseRedBg     = CSI + "101m" // Intense Red background color.
	IntenseGreenBg   = CSI + "102m" // Intense Green background color.
	IntenseYellowBg  = CSI + "103m" // Intense Yellow background color.
	IntenseBlueBg    = CSI + "104m" // Intense Blue background color.
	IntenseMagentaBg = CSI + "105m" // Intense Magenta background color.
	IntenseCyanBg    = CSI + "106m" // Intense Cyan background color.
	IntenseWhiteBg   = CSI + "107m" // Intense White background color.
)

// 8-bit foreground color
// color must be between 0 and 255 otherwise it will return an empty string.
func Foreground8Bit(color int) string {
	if !colorInRange(color) {
		return ""
	}

	return CSI + "38;5;" + strconv.Itoa(color) + "m"
}

// 8-bit background color
// color must be between 0 and 255 otherwise it will return an empty string.
func Background8Bit(color int) string {
	if !colorInRange(color) {
		return ""
	}

	return CSI + "48;5;" + strconv.Itoa(color) + "m"
}

// 24-bit foreground color
// r, g, b must be between 0 and 255 otherwise it will return an empty string.
func Foreground24Bit(r, g, b int) string {
	if !colorInRange(r) || !colorInRange(g) || !colorInRange(b) {
		return ""
	}

	return CSI + "38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

// 24-bit background color
// r, g, b must be between 0 and 255 otherwise it will return an empty string.
func Background24Bit(r, g, b int) string {
	if !colorInRange(r) || !colorInRange(g) || !colorInRange(b) {
		return ""
	}

	return CSI + "48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}
