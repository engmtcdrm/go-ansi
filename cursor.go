package ansi

import "strconv"

const (
	SaveCursorPos    = csi + "s"    // Save the cursor position.
	RestoreCursorPos = csi + "u"    // Restore the cursor position.
	HideCursor       = csi + "?25l" // Hide the cursor.
	ShowCursor       = csi + "?25h" // Show the cursor.
)

var (
	CursorTopLeft      = CursorPosition(1, 1)        // Move the cursor to the top left corner of screen.
	CursorLineBegin    = CursorHorizontalAbsolute(1) // Move the cursor to beginning of the row (line).
	CursorNextLine     = CursorNextLineN(1)          // Move cursor down one row (line).
	CursorPreviousLine = CursorPreviousLineN(1)      // Move cursor up one row (line).
)

// CursorUp moves the cursor up n rows (lines).
// If n is less than 0, it returns an empty string.
func CursorUp(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "A"
}

// CursorDown moves the cursor down n rows (lines).
// If n is less than 0, it returns an empty string.
func CursorDown(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "B"
}

// CursorForward moves the cursor forward n columns.
// If n is less than 0, it returns an empty string.
func CursorForward(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "C"
}

// CursorBackward moves the cursor backward n columns.
// If n is less than 0, it returns an empty string.
func CursorBackward(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "D"
}

// CursorNextLine moves the cursor down n rows (lines).
// If n is less than 0, it returns an empty string.
func CursorNextLineN(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "E"
}

// CursorPreviousLine moves the cursor up n rows (lines).
// If n is less than 0, it returns an empty string.
func CursorPreviousLineN(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "F"
}

// CursorHorizontalAbsolute moves the cursor to the nth column.
// If n is less than 0, it returns an empty string.
func CursorHorizontalAbsolute(n int) string {
	if n < 1 {
		return ""
	}
	return csi + strconv.Itoa(n) + "G"
}

// CursorPosition moves the cursor to the specified position of row (line) and column.
// Both row and column must be greater than or equal to 0, otherwise it returns an empty string.
func CursorPosition(row, column int) string {
	if row < 1 || column < 1 {
		return ""
	}
	return csi + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H"
}
