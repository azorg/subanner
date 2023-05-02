// Font 8x8 compile
// File: "font8make.go"

package font8src

import "fmt"

// ----------------------------------------------------------------------------
func s2b(str string) (retv byte) {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			retv |= 1 << i
		}
	}
	return
}

// ----------------------------------------------------------------------------
func b2s(b []byte) (retv string) {
	for _, b := range b {
		retv += fmt.Sprintf("\\x%02X", b)
	}
	return
}

// ----------------------------------------------------------------------------
func s2s(strs []string) string {
	var b []byte
	for _, s := range strs {
		b = append(b, s2b(s))
	}
	return b2s(b)
}

// ----------------------------------------------------------------------------
func chr2rune(chr string) (retv rune) {
	for _, r := range chr {
		retv = r
		break
	}
	return
}

// ----------------------------------------------------------------------------
const head = `/*
 * Font 8x8
 * File: "font8.go"
 */

package font8

// each symbol has 8 bytes (index by unicode rune)
var Font8 = map[rune]string {
`

// ----------------------------------------------------------------------------
const tail = `} // Font8
`

// ----------------------------------------------------------------------------
func Make() {
	fmt.Print(head)

	for _, sym := range Font8src {
		r := chr2rune(sym.chr)

		chr := sym.chr
		if chr == "\x7F" {
			chr = "[D]"
		} else if chr == `"` {
			chr = "`" + chr + "`"
		} else {
			chr = `"` + chr + `"`
		}

		s := s2s(sym.val)
		fmt.Printf("\t0x%04X: \"%s\", // %s [0x%02X]\n", r, s, chr, sym.cod)
	}

	fmt.Print(tail)
}

// ----------------------------------------------------------------------------

/*** end of "font8make.go" file ***/
