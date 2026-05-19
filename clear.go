package ansi

// Clear and Cursor Control

const (
	ClearFromCursorToEndScreen   = csi + "0J"             // Clears from the cursor to the end of the screen.
	ClearFromCursorToBeginScreen = csi + "1J"             // Clears from the cursor to the beginning of the screen.
	ClearScreen                  = csi + "2J"             // Clear the entire screen.
	ClearToEnd                   = csi + "K"              // Clear from the cursor to the end of the row (line).
	ClearToBegin                 = csi + "1K"             // Clear from the cursor to the beginning of the row (line).
	ClearLine                    = csi + "2K"             // Clear the entire row (line).
	ClearLineReset               = ClearLine + csi + "1G" // Clear the entire row (line) and move the cursor to the start of the row (line).
)
