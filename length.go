package ansi

import "reflect"

const (
	escape          = '\x1b'
	bell            = '\a'
	cancel          = '\x18'
	substitute      = '\x1A'
	stringTerminate = '\x9C'
)

// The logic below is originally from
// https://github.com/clipperhouse/uax29/blob/master/graphemes/ansi.go
//
// It has been modified to reduce cognitive complexity, support runes, and
// conform to terminology in this package, e.g. "input" instead of "data".

// EscapeLength returns the byte length of a valid 7-bit ANSI escape
// sequence at the start of input, or 0 if none.
//
// Recognized forms (ECMA-48 / ISO 6429):
//   - CSI: ESC [ then parameter bytes (0x30-0x3F), intermediate (0x20-0x2F), final (0x40-0x7E)
//   - OSC: ESC ] then payload until BEL (0x07), 7-bit ST (ESC \), CAN (0x18), or SUB (0x1A)
//   - DCS, SOS, PM, APC: ESC P/X/^/_ then payload until 7-bit ST (ESC \), CAN, or SUB
//   - Two-byte: ESC + Fe/Fs (0x40-0x7E excluding above), or Fp (0x30-0x3F), or nF (0x20-0x2F then final)
func EscapeLength[T ~string | ~[]byte | ~[]rune](input T) int {
	rv := reflect.ValueOf(input)
	rt := rv.Type()

	switch rt.Kind() {
	case reflect.String:
		return escapeLength(rv.String())
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			dst := reflect.MakeSlice(reflect.TypeOf([]byte(nil)), rv.Len(), rv.Len())
			reflect.Copy(dst, rv)

			return escapeLength(dst.Interface().([]byte))
		}

		dst := reflect.MakeSlice(reflect.TypeOf([]rune(nil)), rv.Len(), rv.Len())
		reflect.Copy(dst, rv)

		return escapeLengthRune(dst.Interface().([]rune))
	default:
		return 0
	}
}

// escapeLength returns the byte length of a valid 7-bit ANSI escape
// sequence at the start of input, or 0 if none.
//
// Recognized forms (ECMA-48 / ISO 6429):
//   - CSI: ESC [ then parameter bytes (0x30-0x3F), intermediate (0x20-0x2F), final (0x40-0x7E)
//   - OSC: ESC ] then payload until BEL (0x07), 7-bit ST (ESC \), CAN (0x18), or SUB (0x1A)
//   - DCS, SOS, PM, APC: ESC P/X/^/_ then payload until 7-bit ST (ESC \), CAN, or SUB
//   - Two-byte: ESC + Fe/Fs (0x40-0x7E excluding above), or Fp (0x30-0x3F), or nF (0x20-0x2F then final)
func escapeLength[T ~string | ~[]byte](input T) int {
	n := len(input)
	if n < 2 || input[0] != escape {
		return 0
	}

	b1 := input[1]
	switch b1 {
	case '[': // CSI
		body := csiBodyLength(input[2:])
		if body == 0 {
			return 0
		}
		return 2 + body
	case ']': // OSC - allows BEL or 7-bit ST terminator
		body := oscLength(input[2:])
		if body < 0 {
			return 0
		}
		return 2 + body
	case 'P', 'X', '^', '_': // DCS, SOS, PM, APC
		body := stSequenceLength(input[2:])
		if body < 0 {
			return 0
		}
		return 2 + body
	}

	if b1 >= 0x40 && b1 <= 0x7E {
		// Fe/Fs two-byte; [ ] P X ^ _ handled above
		return 2
	}

	if b1 >= 0x30 && b1 <= 0x3F {
		// Fp (private) two-byte
		return 2
	}

	if b1 >= 0x20 && b1 <= 0x2F {
		return intermediateThenFinalLength(input)
	}

	return 0
}

// csiBodyLength returns the length of the CSI body (param/intermediate/final
// bytes). Input is the slice after "ESC [".
//
// Per ECMA-48, the CSI body has the form:
//
//	parameters (0x30–0x3F)*, intermediates (0x20–0x2F)*, final (0x40–0x7E)
//
// Once an intermediate byte is seen, subsequent parameter bytes are invalid.
func csiBodyLength[T ~string | ~[]byte](input T) int {
	seenIntermediate := false
	for i := 0; i < len(input); i++ {
		b := input[i]
		if b >= 0x30 && b <= 0x3F {
			if seenIntermediate {
				return 0
			}
			continue
		}

		if b >= 0x20 && b <= 0x2F {
			seenIntermediate = true
			continue
		}

		if b >= 0x40 && b <= 0x7E {
			return i + 1
		}

		return 0
	}

	return 0
}

// intermediateThenFinalLength returns the length of a two-byte escape sequence
// with intermediates then final. Input is the slice after "ESC".
//
// Per ECMA-48, the sequence has the form:
//
//	intermediate (0x20–0x2F)*, final (0x30–0x7E)
//
// Once an intermediate byte is seen, subsequent parameter bytes are invalid.
func intermediateThenFinalLength[T ~string | ~[]byte](input T) int {
	n := len(input)
	if n < 3 {
		return 0
	}

	i := 2
	for i < n && input[i] >= 0x20 && input[i] <= 0x2F {
		i++
	}

	if i < n && input[i] >= 0x30 && input[i] <= 0x7E {
		return i + 1
	}

	return 0
}

// oscLength returns the length of the OSC body. Input is the slice after
// "ESC ]".
//
// Returns:
//   - n >= 0: consumed body length (includes BEL/ST terminator when present)
//   - -1: not terminated in the provided input
//
// OSC accepts BEL (0x07) or 7-bit ST (ESC \) as terminators by widespread
// convention.
//
// Per ECMA-48, CAN (0x18) and SUB (0x1A) cancel the control string; in that
// case they are not part of the OSC sequence length.
func oscLength[T ~string | ~[]byte](input T) int {
	for i := 0; i < len(input); i++ {
		b := input[i]
		if b == bell {
			return i + 1
		}

		if b == cancel || b == substitute {
			return i
		}

		if b == escape && i+1 < len(input) && input[i+1] == '\\' {
			return i + 2
		}
	}

	return -1
}

// stSequenceLength returns the length of a control-string body. Input is the
// slice after "ESC x".
//
// Returns:
//   - n >= 0: consumed body length (includes ST terminator when present)
//   - -1: not terminated in the provided input
//
// Used for DCS, SOS, PM, and APC, which per ECMA-48 terminate with ST.
// ST here is the 7-bit form (ESC \).
// CAN (0x18) and SUB (0x1A) cancel the control string; in that case they are
// not part of the sequence length.
func stSequenceLength[T ~string | ~[]byte](input T) int {
	for i := 0; i < len(input); i++ {
		if input[i] == cancel || input[i] == substitute {
			return i
		}

		if input[i] == escape && i+1 < len(input) && input[i+1] == '\\' {
			return i + 2
		}
	}

	return -1
}
