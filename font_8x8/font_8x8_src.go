// Font 8x8 source
// -*- encoding: UTF-8 -*-

package main

import "fmt"

var Font = []struct {
	cod int      // index (ASCII/KOI8-R code)
	chr string   // unicode rune
	val []string // [8]
}{
	{0x20, " ", []string{ // "\u0020"
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                "}},

	{0x21, "!", []string{ // "\u0021"
		"                ",
		"      ##        ",
		"      ##        ",
		"      ##        ",
		"      ##        ",
		"      ##        ",
		"                ",
		"      ##        "}},

    {0x22, "\"", []string{ // "\u0022"
    "                ",
    "    ##  ##      ",
    "    ##  ##      ",
    "                ",
    "                ",
    "                ",
    "                ",
    "                "}},

}

func str2byte(str string) uint8 {
	var retv uint8 = 0
	for i := 0; i < 8; i++ {
		if str[2 * i] != ' ' {
			retv |=   1 << i
		}
	}
	return retv
}

const Head = `/*
 * Font 8x8
 * encoding: UTF-8
 */

package font_8x8

// each symbol has 8 bytes (index by unicode rune)
const Font map[uint16] [8]uint8 {
`

func main() {
  fmt.Print(Head)

	for _, sym := range Font {
    fmt.Printf("\t\"%s\": {", sym.chr)
    sep := ""
		for _, str := range sym.val {
      b := str2byte(str)
			fmt.Printf("%s0x%02X", sep, b)
      sep = ", "
		}
    fmt.Printf("} # '%s'=\"\\u%04X\" [0x%02X]\n",
               sym.chr, sym.chr, sym.cod)
	}
}

