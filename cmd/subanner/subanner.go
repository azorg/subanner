/*
 * System-V banner clone with russian cyrillyc support
 * File: "subanner.go"
 */

package main

import (
  "fmt"
  "os"
  "subanner/pkg/font8"
)

func usage() {
  fmt.Println("System-V banner clone with russian cyrillyc support")
  fmt.Println("Usage: subanner [options] word1 [word2...]");
  fmt.Println("Option:");
  fmt.Println("  -h|--help    - show this help");
  fmt.Println("  -d|--dot STR - set dot symbol (like #, @, [])");
}

func main() {
  dot := "#" // by default
  var words []string

  // parse comand line options
  for i := 1; i < len(os.Args); i++ {
    o := os.Args[i]
    switch {
    case o == "-h" || o == "--help":
      usage()
      os.Exit(0)
    case o == "-d" || o == "--dot":
      if i++; i >= len(os.Args) {
        usage()
        os.Exit(-1) // error
      }
      dot = os.Args[i]
    default:
      words = append(words, o)
    }
  }

  // print word(s)
  for _, word := range words {
    font8.Print(word, dot)
  }
}

/*** end of "subanner.go" file ***/

