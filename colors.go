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
	Black   = csi + "30m" // Black foreground color.
	Red     = csi + "31m" // Red foreground color.
	Green   = csi + "32m" // Green foreground color.
	Yellow  = csi + "33m" // Yellow foreground color.
	Blue    = csi + "34m" // Blue foreground color.
	Magenta = csi + "35m" // Magenta foreground color.
	Cyan    = csi + "36m" // Cyan foreground color.
	White   = csi + "37m" // White foreground color.
	Default = csi + "39m" // Default foreground color.
)

// Background Color

const (
	BlackBg   = csi + "40m" // Black background color.
	RedBg     = csi + "41m" // Red background color.
	GreenBg   = csi + "42m" // Green background color.
	YellowBg  = csi + "43m" // Yellow background color.
	BlueBg    = csi + "44m" // Blue background color.
	MagentaBg = csi + "45m" // Magenta background color.
	CyanBg    = csi + "46m" // Cyan background color.
	WhiteBg   = csi + "47m" // White background color.
	DefaultBg = csi + "49m" // Default background color.
)

// Intense Foreground Color

const (
	IntenseBlack   = csi + "90m" // Intense Black foreground color.
	IntenseRed     = csi + "91m" // Intense Red foreground color.
	IntenseGreen   = csi + "92m" // Intense Green foreground color.
	IntenseYellow  = csi + "93m" // Intense Yellow foreground color.
	IntenseBlue    = csi + "94m" // Intense Blue foreground color.
	IntenseMagenta = csi + "95m" // Intense Magenta foreground color.
	IntenseCyan    = csi + "96m" // Intense Cyan foreground color.
	IntenseWhite   = csi + "97m" // Intense White foreground color.

)

// Intense Background Color

const (
	IntenseBlackBg   = csi + "100m" // Intense Black background color.
	IntenseRedBg     = csi + "101m" // Intense Red background color.
	IntenseGreenBg   = csi + "102m" // Intense Green background color.
	IntenseYellowBg  = csi + "103m" // Intense Yellow background color.
	IntenseBlueBg    = csi + "104m" // Intense Blue background color.
	IntenseMagentaBg = csi + "105m" // Intense Magenta background color.
	IntenseCyanBg    = csi + "106m" // Intense Cyan background color.
	IntenseWhiteBg   = csi + "107m" // Intense White background color.
)

// 8-bit foreground color
// color must be between 0 and 255 otherwise it will return an empty string.
func Foreground8Bit(color int) string {
	if !colorInRange(color) {
		return ""
	}

	return csi + "38;5;" + strconv.Itoa(color) + "m"
}

// 8-bit background color
// color must be between 0 and 255 otherwise it will return an empty string.
func Background8Bit(color int) string {
	if !colorInRange(color) {
		return ""
	}

	return csi + "48;5;" + strconv.Itoa(color) + "m"
}

// 24-bit foreground color
// r, g, b must be between 0 and 255 otherwise it will return an empty string.
func Foreground24Bit(r, g, b int) string {
	if !colorInRange(r) || !colorInRange(g) || !colorInRange(b) {
		return ""
	}

	return csi + "38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

// 24-bit background color
// r, g, b must be between 0 and 255 otherwise it will return an empty string.
func Background24Bit(r, g, b int) string {
	if !colorInRange(r) || !colorInRange(g) || !colorInRange(b) {
		return ""
	}

	return csi + "48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}
