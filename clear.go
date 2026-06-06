package ansi

// Clear and Cursor Control

const (
	ClearFromCursorToEndScreen   = CSI + "0J"             // Clears from the cursor to the end of the screen.
	ClearFromCursorToBeginScreen = CSI + "1J"             // Clears from the cursor to the beginning of the screen.
	ClearScreen                  = CSI + "2J"             // Clear the entire screen.
	ClearToEnd                   = CSI + "K"              // Clear from the cursor to the end of the row (line).
	ClearToBegin                 = CSI + "1K"             // Clear from the cursor to the beginning of the row (line).
	ClearLine                    = CSI + "2K"             // Clear the entire row (line).
	ClearLineReset               = ClearLine + CSI + "1G" // Clear the entire row (line) and move the cursor to the start of the row (line).
)
