/*
 * Font 8x8: print word
 * File: "font8print.go"
 */

package font8

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Print big string (word) to console
// (set dot="#" or "[]" for example)
func Print(word, dot string) {
	if len(dot) == 0 {
		dot = "#" // set default
	}
	off := strings.Repeat(" ", utf8.RuneCountInString(dot))

	for i := 0; i < 8; i++ {
		for _, ch := range word {
			s, ok := Font8[ch]
			if !ok {
				continue // unknown symbol
			}
			b := s[i]
			for j := 0; j < 8; j++ {
				c := off
				if (b & (1 << j)) != 0 {
					c = dot
				}
				//c := (b & (1 << j)) != 0 ? on : off
				fmt.Print(c)
			}
		}
		fmt.Println()
	}
}

/*** end of "font8print.go" file ***/
