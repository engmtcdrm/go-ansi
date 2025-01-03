# ansi

```go
import "github.com/engmtcdrm/go-ansi"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func Background24Bit\(r, g, b int\) string](<#func-background24bit>)
- [func Background8Bit\(color int\) string](<#func-background8bit>)
- [func CursorBackward\(n int\) string](<#func-cursorbackward>)
- [func CursorDown\(n int\) string](<#func-cursordown>)
- [func CursorForward\(n int\) string](<#func-cursorforward>)
- [func CursorHorizontalAbsolute\(n int\) string](<#func-cursorhorizontalabsolute>)
- [func CursorNextLineN\(n int\) string](<#func-cursornextlinen>)
- [func CursorPosition\(row, column int\) string](<#func-cursorposition>)
- [func CursorPreviousLineN\(n int\) string](<#func-cursorpreviouslinen>)
- [func CursorUp\(n int\) string](<#func-cursorup>)
- [func Foreground24Bit\(r, g, b int\) string](<#func-foreground24bit>)
- [func Foreground8Bit\(color int\) string](<#func-foreground8bit>)
- [func ScrollDownN\(n int\) string](<#func-scrolldownn>)
- [func ScrollUpN\(n int\) string](<#func-scrollupn>)
- [func StripCodes\(input string\) string](<#func-stripcodes>)

## Constants

```go
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
```

```go
const (

    // Clears from the cursor to the end of the screen.
    ClearFromCursorToEndScreen = csi + "0J"
    // Clears from the cursor to the beginning of the screen.
    ClearFromCursorToBeginScreen = csi + "1J"
    // Clear the entire screen.
    ClearScreen = csi + "2J"
    // Clear from the cursor to the end of the row (line).
    ClearToEnd = csi + "K"
    // Clear from the cursor to the beginning of the row (line).
    ClearToBegin = csi + "1K"
    // Clear the entire row (line).
    ClearLine = csi + "2K"
    // Clear the entire row (line) and move the cursor to the start of the row (line).
    ClearLineReset = csi + "2k\r"

    // Save the cursor position
    SaveCursorPos = csi + "s"
    // Restore the cursor position
    RestoreCursorPos = csi + "u"

    // Hide the cursor
    HideCursor = csi + "?25l"

    // Show the cursor
    ShowCursor = csi + "?25h"

    // Text Formatting
    Reset           = csi + "0m"
    Bold            = csi + "1m"
    Dim             = csi + "2m"
    Italic          = csi + "3m"
    Underline       = csi + "4m"
    SlowBlink       = csi + "5m"
    RapidBlink      = csi + "6m"
    Reverse         = csi + "7m"
    Hidden          = csi + "8m"
    Strike          = csi + "9m"
    DoubleUnderline = csi + "21m"

    // Reset Formatting
    ResetIntensity = csi + "22m"
    ResetItalic    = csi + "23m"
    ResetUnderline = csi + "24m"
    ResetBlink     = csi + "25m"
    ResetReverse   = csi + "27m"
    ResetHidden    = csi + "28m"
    ResetStrike    = csi + "29m"

    // Foreground Color
    Black   = csi + "30m"
    Red     = csi + "31m"
    Green   = csi + "32m"
    Yellow  = csi + "33m"
    Blue    = csi + "34m"
    Magenta = csi + "35m"
    Cyan    = csi + "36m"
    White   = csi + "37m"

    // Background Color
    BlackBg   = csi + "40m"
    RedBg     = csi + "41m"
    GreenBg   = csi + "42m"
    YellowBg  = csi + "43m"
    BlueBg    = csi + "44m"
    MagentaBg = csi + "45m"
    CyanBg    = csi + "46m"
    WhiteBg   = csi + "47m"
    DefaultBg = csi + "49m"

    // Intense Foreground Color
    IntenseBlack   = csi + "90m"
    IntenseRed     = csi + "91m"
    IntenseGreen   = csi + "92m"
    IntenseYellow  = csi + "93m"
    IntenseBlue    = csi + "94m"
    IntenseMagenta = csi + "95m"
    IntenseCyan    = csi + "96m"
    IntenseWhite   = csi + "97m"

    // Intense Background Color
    IntenseBlackBg   = csi + "100m"
    IntenseRedBg     = csi + "101m"
    IntenseGreenBg   = csi + "102m"
    IntenseYellowBg  = csi + "103m"
    IntenseBlueBg    = csi + "104m"
    IntenseMagentaBg = csi + "105m"
    IntenseCyanBg    = csi + "106m"
    IntenseWhiteBg   = csi + "107m"
)
```

## Variables

```go
var (
    // Move the cursor to the top left corner of screen
    CursorTopLeft = CursorPosition(1, 1)
    // Move the cursor to beginning of the row (line).
    CursorLineBegin = CursorHorizontalAbsolute(1)
    // Move cursor down one row (line).
    CursorNextLine = CursorNextLineN(1)
    // Move cursor up one row (line).
    CursorPreviousLine = CursorPreviousLineN(1)

    // Scroll the screen up one row (line).
    ScrollUp1 = ScrollUpN(1)
    // Scroll the screen down one row (line).
    ScrollDown1 = ScrollDownN(1)
)
```

## func [Background24Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L174>)

```go
func Background24Bit(r, g, b int) string
```

24\-bit background color r, g, b must be between 0 and 255 otherwise it will return an empty string

## func [Background8Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L154>)

```go
func Background8Bit(color int) string
```

8\-bit background color color must be between 0 and 255 othrewise it will return an empty string

## func [CursorBackward](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L198>)

```go
func CursorBackward(n int) string
```

CursorBackward moves the cursor backward n columns.

## func [CursorDown](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L188>)

```go
func CursorDown(n int) string
```

CursorDown moves the cursor down n rows \(lines\).

## func [CursorForward](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L193>)

```go
func CursorForward(n int) string
```

CursorForward moves the cursor forward n columns.

## func [CursorHorizontalAbsolute](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L213>)

```go
func CursorHorizontalAbsolute(n int) string
```

CursorHorizontalAbsolute moves the cursor to the nth column.

## func [CursorNextLineN](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L203>)

```go
func CursorNextLineN(n int) string
```

CursorNextLine moves the cursor down n rows \(lines\).

## func [CursorPosition](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L218>)

```go
func CursorPosition(row, column int) string
```

CursorPosition moves the cursor to the specified position of row \(line\) and column.

## func [CursorPreviousLineN](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L208>)

```go
func CursorPreviousLineN(n int) string
```

CursorPreviousLine moves the cursor up n rows \(lines\).

## func [CursorUp](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L183>)

```go
func CursorUp(n int) string
```

CursorUp moves the cursor up n rows \(lines\).

## func [Foreground24Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L164>)

```go
func Foreground24Bit(r, g, b int) string
```

24\-bit foreground color r, g, b must be between 0 and 255 otherwise it will return an empty string

## func [Foreground8Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L144>)

```go
func Foreground8Bit(color int) string
```

8\-bit foreground color color must be between 0 and 255 otherwise it will return an empty string

## func [ScrollDownN](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L228>)

```go
func ScrollDownN(n int) string
```

ScrollDownN scrolls the screen down n rows \(lines\).

## func [ScrollUpN](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L223>)

```go
func ScrollUpN(n int) string
```

ScrollUp scrolls the screen up n rows \(lines\).

## func [StripCodes](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L233>)

```go
func StripCodes(input string) string
```

StripCodes removes all ANSI escape codes from the input string.
