## Description

The `ansi` package provides a set of constants and functions for generating ANSI escape codes, which are used to control text formatting, color, and other output options on text terminals. These escape codes can be used to manipulate the appearance of text in terminal emulators, such as changing text color, moving the cursor, clearing parts of the screen, and more.

Key Features:

- Text Formatting: Apply styles like bold, italic, underline, blink, reverse, hidden, and strike-through to text.
- Color Control: Set foreground and background colors using standard and intense colors.
- Cursor Movement: Move the cursor to specific positions, save and restore cursor positions, and hide or show the cursor.
- Screen Clearing: Clear parts of the screen or entire lines.
- Scrolling: Scroll the screen up or down by a specified number of lines.
- 8-bit Color Support: Set foreground and background colors using 8-bit color codes.
- 24-bit Color Support: Set foreground and background colors using 24-bit color codes.

## Installation

ansi is available using the standard `go get` command.

Install by running:

```sh
go get github.com/engmtcdrm/go-ansi
```

Run tests by running:

```sh
go test github.com/engmtcdrm/go-ansi
```

## Usage

```go
import "github.com/engmtcdrm/go-ansi"
```

## Index

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Index](#index)
- [Constants](#constants)
- [Variables](#variables)
- [func Background24Bit](#func-background24bit)
- [func Background8Bit](#func-background8bit)
- [func CursorBackward](#func-cursorbackward)
- [func CursorDown](#func-cursordown)
- [func CursorForward](#func-cursorforward)
- [func CursorHorizontalAbsolute](#func-cursorhorizontalabsolute)
- [func CursorNextLineN](#func-cursornextlinen)
- [func CursorPosition](#func-cursorposition)
- [func CursorPreviousLineN](#func-cursorpreviouslinen)
- [func CursorUp](#func-cursorup)
- [func Foreground24Bit](#func-foreground24bit)
- [func Foreground8Bit](#func-foreground8bit)
- [func ScrollDownN](#func-scrolldownn)
- [func ScrollUpN](#func-scrollupn)
- [func Strip](#func-strip)
- [func StripCodes](#func-stripcodes)


## Constants

<a name="ClearFromCursorToEndScreen"></a>

```go
const (
    ClearFromCursorToEndScreen   = csi + "0J" // Clears from the cursor to the end of the screen.
    ClearFromCursorToBeginScreen = csi + "1J" // Clears from the cursor to the beginning of the screen.
    ClearScreen                  = csi + "2J" // Clear the entire screen.
    ClearToEnd                   = csi + "K"  // Clear from the cursor to the end of the row (line).
    ClearToBegin                 = csi + "1K" // Clear from the cursor to the beginning of the row (line).
    ClearLine                    = csi + "2K" // Clear the entire row (line).
)
```

<a name="ANSIBlack"></a>ANSI color codes.

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

<a name="Black"></a>

```go
const (
    Black   = csi + "30m" // Black foreground color.
    Red     = csi + "31m" // Red foreground color.
    Green   = csi + "32m" // Green foreground color.
    Yellow  = csi + "33m" // Yellow foreground color.
    Blue    = csi + "34m" // Blue foreground color.
    Magenta = csi + "35m" // Magenta foreground color.
    Cyan    = csi + "36m" // Cyan foreground color.
    White   = csi + "37m" // White foreground color.
)
```

<a name="BlackBg"></a>

```go
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
```

<a name="IntenseBlack"></a>

```go
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
```

<a name="IntenseBlackBg"></a>

```go
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
```

<a name="SaveCursorPos"></a>

```go
const (
    SaveCursorPos    = csi + "s"    // Save the cursor position.
    RestoreCursorPos = csi + "u"    // Restore the cursor position.
    HideCursor       = csi + "?25l" // Hide the cursor.
    ShowCursor       = csi + "?25h" // Show the cursor.
)
```

<a name="Reset"></a>

```go
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

    ResetIntensity = csi + "22m" // Reset bold or faint.
    ResetItalic    = csi + "23m" // Reset italic.
    ResetUnderline = csi + "24m" // Reset underline.
    ResetBlink     = csi + "25m" // Reset blink.
    ResetReverse   = csi + "27m" // Reset reverse.
    ResetHidden    = csi + "28m" // Reset hidden.
    ResetStrike    = csi + "29m" // Reset strikethrough.
)
```

## Variables

<a name="CursorTopLeft"></a>

```go
var (
    CursorTopLeft      = CursorPosition(1, 1)        // Move the cursor to the top left corner of screen.
    CursorLineBegin    = CursorHorizontalAbsolute(1) // Move the cursor to beginning of the row (line).
    CursorNextLine     = CursorNextLineN(1)          // Move cursor down one row (line).
    CursorPreviousLine = CursorPreviousLineN(1)      // Move cursor up one row (line).
)
```

<a name="ScrollUp1"></a>

```go
var (
    ScrollUp1   = ScrollUpN(1)   // Scroll the screen up one row (line).
    ScrollDown1 = ScrollDownN(1) // Scroll the screen down one row (line).
)
```

<a name="ClearLineReset"></a>

```go
var (
    ClearLineReset = ClearLine + CursorLineBegin // Clear the entire row (line) and move the cursor to the start of the row (line).
)
```

<a name="Background24Bit"></a>
## func [Background24Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/colors.go#L108>)

```go
func Background24Bit(r, g, b int) string
```

24\-bit background color r, g, b must be between 0 and 255 otherwise it will return an empty string.

<a name="Background8Bit"></a>
## func [Background8Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/colors.go#L90>)

```go
func Background8Bit(color int) string
```

8\-bit background color color must be between 0 and 255 otherwise it will return an empty string.

<a name="CursorBackward"></a>
## func [CursorBackward](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L48>)

```go
func CursorBackward(n int) string
```

CursorBackward moves the cursor backward n columns. If n is negative, it returns an empty string.

<a name="CursorDown"></a>
## func [CursorDown](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L30>)

```go
func CursorDown(n int) string
```

CursorDown moves the cursor down n rows \(lines\). If n is negative, it returns an empty string.

<a name="CursorForward"></a>
## func [CursorForward](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L39>)

```go
func CursorForward(n int) string
```

CursorForward moves the cursor forward n columns. If n is negative, it returns an empty string.

<a name="CursorHorizontalAbsolute"></a>
## func [CursorHorizontalAbsolute](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L75>)

```go
func CursorHorizontalAbsolute(n int) string
```

CursorHorizontalAbsolute moves the cursor to the nth column. If n is negative, it returns an empty string.

<a name="CursorNextLineN"></a>
## func [CursorNextLineN](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L57>)

```go
func CursorNextLineN(n int) string
```

CursorNextLine moves the cursor down n rows \(lines\). If n is negative, it returns an empty string.

<a name="CursorPosition"></a>
## func [CursorPosition](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L84>)

```go
func CursorPosition(row, column int) string
```

CursorPosition moves the cursor to the specified position of row \(line\) and column. Both row and column must be greater than or equal to 0, otherwise it returns an empty string.

<a name="CursorPreviousLineN"></a>
## func [CursorPreviousLineN](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L66>)

```go
func CursorPreviousLineN(n int) string
```

CursorPreviousLine moves the cursor up n rows \(lines\). If n is negative, it returns an empty string.

<a name="CursorUp"></a>
## func [CursorUp](<https://github.com/engmtcdrm/go-ansi/blob/main/cursor.go#L21>)

```go
func CursorUp(n int) string
```

CursorUp moves the cursor up n rows \(lines\). If n is negative, it returns an empty string.

<a name="Foreground24Bit"></a>
## func [Foreground24Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/colors.go#L99>)

```go
func Foreground24Bit(r, g, b int) string
```

24\-bit foreground color r, g, b must be between 0 and 255 otherwise it will return an empty string.

<a name="Foreground8Bit"></a>
## func [Foreground8Bit](<https://github.com/engmtcdrm/go-ansi/blob/main/colors.go#L81>)

```go
func Foreground8Bit(color int) string
```

8\-bit foreground color color must be between 0 and 255 otherwise it will return an empty string.

<a name="ScrollDownN"></a>
## func [ScrollDownN](<https://github.com/engmtcdrm/go-ansi/blob/main/scroll.go#L21>)

```go
func ScrollDownN(n int) string
```

ScrollDownN scrolls the screen down n rows \(lines\). If n is negative, it returns an empty string.

<a name="ScrollUpN"></a>
## func [ScrollUpN](<https://github.com/engmtcdrm/go-ansi/blob/main/scroll.go#L12>)

```go
func ScrollUpN(n int) string
```

ScrollUp scrolls the screen up n rows \(lines\). If n is negative, it returns an empty string.

<a name="Strip"></a>
## func [Strip](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L11>)

```go
func Strip(input string) string
```

Strip removes all ANSI escape codes from the input string.

<a name="StripCodes"></a>
## func [StripCodes](<https://github.com/engmtcdrm/go-ansi/blob/main/ansi.go#L18>)

```go
func StripCodes(input string) string
```

StripCodes removes all ANSI escape codes from the input string.

Deprecated: Use [Strip](<#Strip>) instead. This will be removed in v2.0.0.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
