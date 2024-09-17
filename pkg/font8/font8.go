/*
 * Font 8x8
 * File: "font8.go"
 */

package font8

// each symbol has 8 bytes (index by unicode rune)
var Font8 = map[rune]string{
	0x0020: "\x00\x00\x00\x00\x00\x00\x00\x00", // " " [0x20]
	0x0021: "\x08\x1C\x1C\x08\x08\x00\x08\x00", // "!" [0x21]
	0x0022: "\x14\x14\x00\x00\x00\x00\x00\x00", // `"` [0x22]
	0x0023: "\x14\x14\x7F\x14\x7F\x14\x14\x00", // "#" [0x23]
	0x0024: "\x3E\x49\x09\x3E\x48\x49\x3E\x00", // "$" [0x24]
	0x0025: "\x47\x25\x17\x08\x74\x52\x71\x00", // "%" [0x25]
	0x0026: "\x0C\x12\x0C\x0E\x51\x21\x5E\x00", // "&" [0x26]
	0x0027: "\x08\x08\x00\x00\x00\x00\x00\x00", // "'" [0x27]
	0x0028: "\x18\x04\x02\x02\x02\x04\x18\x00", // "(" [0x28]
	0x0029: "\x0C\x10\x20\x20\x20\x10\x0C\x00", // ")" [0x29]
	0x002A: "\x00\x22\x14\x77\x14\x22\x00\x00", // "*" [0x2A]
	0x002B: "\x00\x08\x08\x3E\x08\x08\x00\x00", // "+" [0x2B]
	0x002C: "\x00\x00\x00\x00\x0C\x08\x04\x00", // "," [0x2C]
	0x002D: "\x00\x00\x00\x3E\x00\x00\x00\x00", // "-" [0x2D]
	0x002E: "\x00\x00\x00\x00\x00\x18\x18\x00", // "." [0x2E]
	0x002F: "\x40\x20\x10\x08\x04\x02\x01\x00", // "/" [0x2F]
	0x0030: "\x3E\x61\x51\x49\x45\x43\x3E\x00", // "0" [0x30]
	0x0031: "\x08\x0C\x0A\x08\x08\x08\x3E\x00", // "1" [0x31]
	0x0032: "\x3E\x41\x40\x3E\x01\x01\x7F\x00", // "2" [0x32]
	0x0033: "\x3E\x41\x40\x3E\x40\x41\x3E\x00", // "3" [0x33]
	0x0034: "\x30\x28\x24\x22\x7F\x20\x20\x00", // "4" [0x34]
	0x0035: "\x7F\x01\x01\x3E\x40\x41\x3E\x00", // "5" [0x35]
	0x0036: "\x3E\x41\x01\x3F\x41\x41\x3E\x00", // "6" [0x36]
	0x0037: "\x7F\x21\x10\x08\x04\x04\x04\x00", // "7" [0x37]
	0x0038: "\x3E\x41\x41\x3E\x41\x41\x3E\x00", // "8" [0x38]
	0x0039: "\x3E\x41\x41\x7E\x40\x41\x3E\x00", // "9" [0x39]
	0x003A: "\x00\x18\x18\x00\x18\x18\x00\x00", // ":" [0x3A]
	0x003B: "\x00\x18\x18\x00\x18\x18\x08\x00", // ";" [0x3B]
	0x003C: "\x10\x08\x04\x02\x04\x08\x10\x00", // "<" [0x3C]
	0x003D: "\x00\x00\x3E\x00\x3E\x00\x00\x00", // "=" [0x3D]
	0x003E: "\x04\x08\x10\x20\x10\x08\x04\x00", // ">" [0x3E]
	0x003F: "\x3E\x41\x40\x30\x08\x00\x08\x00", // "?" [0x3F]
	0x0040: "\x3E\x41\x79\x45\x39\x01\x3E\x00", // "@" [0x40]
	0x0041: "\x1C\x22\x41\x7F\x41\x41\x41\x00", // "A" [0x41]
	0x0042: "\x3F\x41\x41\x3F\x41\x41\x3F\x00", // "B" [0x42]
	0x0043: "\x3E\x41\x01\x01\x01\x41\x3E\x00", // "C" [0x43]
	0x0044: "\x1F\x21\x41\x41\x41\x41\x3F\x00", // "D" [0x44]
	0x0045: "\x7F\x01\x01\x1F\x01\x01\x7F\x00", // "E" [0x45]
	0x0046: "\x7F\x01\x01\x1F\x01\x01\x01\x00", // "F" [0x46]
	0x0047: "\x3E\x41\x01\x79\x41\x41\x3E\x00", // "G" [0x47]
	0x0048: "\x41\x41\x41\x7F\x41\x41\x41\x00", // "H" [0x48]
	0x0049: "\x1C\x08\x08\x08\x08\x08\x1C\x00", // "I" [0x49]
	0x004A: "\x70\x20\x20\x20\x20\x21\x1E\x00", // "J" [0x4A]
	0x004B: "\x21\x11\x09\x07\x09\x11\x21\x00", // "K" [0x4B]
	0x004C: "\x01\x01\x01\x01\x01\x01\x7F\x00", // "L" [0x4C]
	0x004D: "\x41\x63\x55\x49\x41\x41\x41\x00", // "M" [0x4D]
	0x004E: "\x41\x43\x45\x49\x51\x61\x41\x00", // "N" [0x4E]
	0x004F: "\x3E\x41\x41\x41\x41\x41\x3E\x00", // "O" [0x4F]
	0x0050: "\x3F\x41\x41\x3F\x01\x01\x01\x00", // "P" [0x50]
	0x0051: "\x3E\x41\x41\x41\x51\x21\x5E\x00", // "Q" [0x51]
	0x0052: "\x3F\x41\x41\x3F\x11\x21\x41\x00", // "R" [0x52]
	0x0053: "\x3E\x41\x01\x3E\x40\x41\x3E\x00", // "S" [0x53]
	0x0054: "\x7F\x08\x08\x08\x08\x08\x08\x00", // "T" [0x54]
	0x0055: "\x41\x41\x41\x41\x41\x41\x3E\x00", // "U" [0x55]
	0x0056: "\x41\x41\x41\x41\x22\x14\x08\x00", // "V" [0x56]
	0x0057: "\x41\x41\x49\x49\x49\x49\x36\x00", // "W" [0x57]
	0x0058: "\x41\x22\x14\x08\x14\x22\x41\x00", // "X" [0x58]
	0x0059: "\x41\x22\x14\x08\x08\x08\x08\x00", // "Y" [0x59]
	0x005A: "\x7F\x20\x10\x08\x04\x02\x7F\x00", // "Z" [0x5A]
	0x005B: "\x3E\x02\x02\x02\x02\x02\x3E\x00", // "[" [0x5B]
	0x005C: "\x01\x02\x04\x08\x10\x20\x40\x00", // "\" [0x5C]
	0x005D: "\x3E\x20\x20\x20\x20\x20\x3E\x00", // "]" [0x5D]
	0x005E: "\x08\x14\x22\x00\x00\x00\x00\x00", // "^" [0x5E]
	0x005F: "\x00\x00\x00\x00\x00\x00\x7F\x00", // "_" [0x5F]
	0x0060: "\x0C\x30\x00\x00\x00\x00\x00\x00", // "`" [0x60]
	0x0061: "\x00\x00\x3E\x21\x21\x21\x5E\x00", // "a" [0x61]
	0x0062: "\x01\x01\x3F\x41\x41\x41\x3F\x00", // "b" [0x62]
	0x0063: "\x00\x00\x3E\x41\x01\x41\x3E\x00", // "c" [0x63]
	0x0064: "\x20\x20\x3E\x21\x21\x21\x5E\x00", // "d" [0x64]
	0x0065: "\x00\x00\x3E\x41\x7F\x01\x3E\x00", // "e" [0x65]
	0x0066: "\x18\x24\x04\x0E\x04\x04\x04\x00", // "f" [0x66]
	0x0067: "\x00\x00\x3E\x41\x7E\x40\x3E\x00", // "g" [0x67]
	0x0068: "\x02\x02\x3A\x46\x42\x42\x42\x00", // "h" [0x68]
	0x0069: "\x08\x00\x0C\x08\x08\x08\x1C\x00", // "i" [0x69]
	0x006A: "\x10\x00\x18\x10\x10\x12\x0C\x00", // "j" [0x6A]
	0x006B: "\x02\x02\x22\x12\x0E\x12\x22\x00", // "k" [0x6B]
	0x006C: "\x0C\x08\x08\x08\x08\x08\x1C\x00", // "l" [0x6C]
	0x006D: "\x00\x00\x37\x49\x49\x49\x49\x00", // "m" [0x6D]
	0x006E: "\x00\x00\x3A\x46\x42\x42\x42\x00", // "n" [0x6E]
	0x006F: "\x00\x00\x3E\x41\x41\x41\x3E\x00", // "o" [0x6F]
	0x0070: "\x00\x00\x3E\x42\x42\x3E\x02\x00", // "p" [0x70]
	0x0071: "\x00\x00\x7C\x42\x42\x7C\x40\x00", // "q" [0x71]
	0x0072: "\x00\x00\x3A\x46\x02\x02\x02\x00", // "r" [0x72]
	0x0073: "\x00\x00\x3C\x02\x3C\x40\x3E\x00", // "s" [0x73]
	0x0074: "\x04\x04\x0E\x04\x04\x24\x18\x00", // "t" [0x74]
	0x0075: "\x00\x00\x42\x42\x42\x62\x5C\x00", // "u" [0x75]
	0x0076: "\x00\x00\x41\x41\x22\x14\x08\x00", // "v" [0x76]
	0x0077: "\x00\x00\x41\x41\x49\x49\x36\x00", // "w" [0x77]
	0x0078: "\x00\x00\x22\x14\x08\x14\x22\x00", // "x" [0x78]
	0x0079: "\x00\x00\x42\x42\x7C\x40\x3C\x00", // "y" [0x79]
	0x007A: "\x00\x00\x3E\x10\x08\x04\x3E\x00", // "z" [0x7A]
	0x007B: "\x38\x04\x04\x06\x04\x04\x38\x00", // "{" [0x7B]
	0x007C: "\x08\x08\x08\x08\x08\x08\x08\x00", // "|" [0x7C]
	0x007D: "\x0E\x10\x10\x30\x10\x10\x0E\x00", // "}" [0x7D]
	0x007E: "\x00\x00\x06\x49\x30\x00\x00\x00", // "~" [0x7E]
	0x007F: "\x55\xAA\x55\xAA\x55\xAA\x55\x00", // [D] [0x7F]
	0x2500: "\x00\x00\x00\xFF\x00\x00\x00\x00", // "─" [0x80]
	0x2502: "\x08\x08\x08\x08\x08\x08\x08\x08", // "│" [0x81]
	0x250C: "\x00\x00\x00\xF8\x08\x08\x08\x08", // "┌" [0x82]
	0x2510: "\x00\x00\x00\x0F\x08\x08\x08\x08", // "┐" [0x83]
	0x2514: "\x08\x08\x08\xF8\x00\x00\x00\x00", // "└" [0x84]
	0x2518: "\x08\x08\x08\x0F\x00\x00\x00\x00", // "┘" [0x85]
	0x251C: "\x08\x08\x08\xF8\x08\x08\x08\x08", // "├" [0x86]
	0x2524: "\x08\x08\x08\x0F\x08\x08\x08\x08", // "┤" [0x87]
	0x252C: "\x00\x00\x00\xFF\x08\x08\x08\x08", // "┬" [0x88]
	0x2534: "\x08\x08\x08\xFF\x00\x00\x00\x00", // "┴" [0x89]
	0x253C: "\x08\x08\x08\xFF\x08\x08\x08\x08", // "┼" [0x8A]
	0x2580: "\xFF\xFF\xFF\xFF\x00\x00\x00\x00", // "▀" [0x8B]
	0x2584: "\x00\x00\x00\x00\xFF\xFF\xFF\xFF", // "▄" [0x8C]
	0x2588: "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF", // "█" [0x8D]
	0x258C: "\x0F\x0F\x0F\x0F\x0F\x0F\x0F\x0F", // "▌" [0x8E]
	0x2590: "\xF0\xF0\xF0\xF0\xF0\xF0\xF0\xF0", // "▐" [0x8F]
	0x2591: "\x49\x00\x92\x00\x49\x00\x92\x00", // "░" [0x90]
	0x2592: "\x49\x92\x24\x49\x92\x24\x49\x92", // "▒" [0x91]
	0x2593: "\x55\xAA\x55\xAA\x55\xAA\x55\xAA", // "▓" [0x92]
	0x2320: "\x30\x48\x08\x08\x08\x08\x08\x08", // "⌠" [0x93]
	0x25A0: "\x00\x7F\x7F\x7F\x7F\x7F\x00\x00", // "■" [0x94]
	0x2219: "\x00\x00\x18\x3C\x18\x00\x00\x00", // "∙" [0x95]
	0x221A: "\x70\x10\x11\x12\x14\x18\x10\x00", // "√" [0x96]
	0x2248: "\x00\x4C\x32\x00\x4C\x32\x00\x00", // "≈" [0x97]
	0x2264: "\x00\x08\x04\x02\x04\x08\x00\x0E", // "≤" [0x98]
	0x2265: "\x00\x02\x04\x08\x04\x02\x00\x0E", // "≥" [0x99]
	0x00A0: "\x00\x00\x00\x00\x00\x00\x00\x00", // " " [0x9A]
	0x2321: "\x08\x08\x08\x08\x08\x08\x09\x06", // "⌡" [0x9B]
	0x00B0: "\x18\x24\x24\x18\x00\x00\x00\x00", // "°" [0x9C]
	0x00B2: "\x38\x44\x30\x08\x7C\x00\x00\x00", // "²" [0x9D]
	0x00B7: "\x00\x00\x00\x08\x00\x00\x00\x00", // "·" [0x9E]
	0x00F7: "\x00\x08\x00\x3E\x00\x08\x00\x00", // "÷" [0x9F]
	0x2550: "\x00\x00\xFF\x00\xFF\x00\x00\x00", // "═" [0xA0]
	0x2551: "\x14\x14\x14\x14\x14\x14\x14\x14", // "║" [0xA1]
	0x2552: "\x00\x00\xF8\x08\xF8\x08\x08\x08", // "╒" [0xA2]
	0x0451: "\x14\x00\x3E\x41\x7F\x01\x3E\x00", // "ё" [0xA3]
	0x2553: "\x00\x00\x00\xFC\x14\x14\x14\x14", // "╓" [0xA4]
	0x2554: "\x00\x00\xFC\x04\xF4\x14\x14\x14", // "╔" [0xA5]
	0x2555: "\x00\x00\x0F\x08\x0F\x08\x08\x08", // "╕" [0xA6]
	0x2556: "\x00\x00\x00\x1F\x14\x14\x14\x14", // "╖" [0xA7]
	0x2557: "\x00\x00\x1F\x10\x17\x14\x14\x14", // "╗" [0xA8]
	0x2558: "\x08\x08\xF8\x08\xF8\x00\x00\x00", // "╘" [0xA9]
	0x2559: "\x14\x14\x14\xFC\x00\x00\x00\x00", // "╙" [0xAA]
	0x255A: "\x14\x14\xF4\x04\xFC\x00\x00\x00", // "╚" [0xAB]
	0x255B: "\x08\x08\x0F\x08\x0F\x00\x00\x00", // "╛" [0xAC]
	0x255C: "\x14\x14\x14\x1F\x00\x00\x00\x00", // "╜" [0xAD]
	0x255D: "\x14\x14\x17\x10\x1F\x00\x00\x00", // "╝" [0xAE]
	0x255E: "\x08\x08\xF8\x08\xF8\x08\x08\x08", // "╞" [0xAF]
	0x255F: "\x14\x14\x14\xF4\x14\x14\x14\x14", // "╟" [0xB0]
	0x2560: "\x14\x14\xF4\x04\xF4\x14\x14\x14", // "╠" [0xB1]
	0x2561: "\x08\x08\x0F\x08\x0F\x08\x08\x08", // "╡" [0xB2]
	0x0401: "\x14\x7F\x01\x1F\x01\x01\x7F\x00", // "Ё" [0xB3]
	0x2562: "\x14\x14\x14\x17\x14\x14\x14\x14", // "╢" [0xB4]
	0x2563: "\x14\x14\x17\x10\x17\x14\x14\x14", // "╣" [0xB5]
	0x2564: "\x00\x00\xFF\x00\xFF\x08\x08\x08", // "╤" [0xB6]
	0x2565: "\x00\x00\x00\xFF\x14\x14\x14\x14", // "╥" [0xB7]
	0x2566: "\x00\x00\xFF\x00\xF7\x14\x14\x14", // "╦" [0xB8]
	0x2567: "\x08\x08\xFF\x00\xFF\x00\x00\x00", // "╧" [0xB9]
	0x2568: "\x14\x14\x14\xFF\x00\x00\x00\x00", // "╨" [0xBA]
	0x2569: "\x14\x14\xF7\x00\xFF\x00\x00\x00", // "╩" [0xBB]
	0x256A: "\x08\x08\xFF\x00\xFF\x08\x08\x08", // "╪" [0xBC]
	0x256B: "\x14\x14\x14\xFF\x14\x14\x14\x14", // "╫" [0xBD]
	0x256C: "\x14\x14\xF7\x00\xF7\x14\x14\x14", // "╬" [0xBE]
	0x00A9: "\x7E\x81\xBD\x85\xBD\x81\x7E\x00", // "©" [0xBF]
	0x044E: "\x00\x00\x31\x49\x4F\x49\x31\x00", // "ю" [0xC0]
	0x0430: "\x00\x00\x3E\x21\x21\x21\x5E\x00", // "а" [0xC1]
	0x0431: "\x3E\x01\x3D\x43\x41\x41\x3F\x00", // "б" [0xC2]
	0x0446: "\x00\x00\x21\x21\x21\x7F\x40\x00", // "ц" [0xC3]
	0x0434: "\x00\x00\x3C\x22\x22\x7F\x41\x00", // "д" [0xC4]
	0x0435: "\x00\x00\x3E\x41\x7F\x01\x3E\x00", // "е" [0xC5]
	0x0444: "\x00\x08\x3E\x49\x3E\x08\x08\x00", // "ф" [0xC6]
	0x0433: "\x00\x00\x3E\x02\x02\x02\x02\x00", // "г" [0xC7]
	0x0445: "\x00\x00\x22\x14\x08\x14\x22\x00", // "х" [0xC8]
	0x0438: "\x00\x00\x42\x42\x42\x62\x5C\x00", // "и" [0xC9]
	0x0439: "\x24\x18\x42\x42\x42\x62\x5C\x00", // "й" [0xCA]
	0x043A: "\x00\x00\x32\x0A\x0E\x12\x22\x00", // "к" [0xCB]
	0x043B: "\x00\x00\x78\x44\x44\x44\x42\x00", // "л" [0xCC]
	0x043C: "\x00\x00\x41\x63\x55\x49\x41\x00", // "м" [0xCD]
	0x043D: "\x00\x00\x41\x41\x7F\x41\x41\x00", // "н" [0xCE]
	0x043E: "\x00\x00\x3E\x41\x41\x41\x3E\x00", // "о" [0xCF]
	0x043F: "\x00\x00\x7F\x41\x41\x41\x41\x00", // "п" [0xD0]
	0x044F: "\x00\x00\x7E\x41\x7E\x44\x43\x00", // "я" [0xD1]
	0x0440: "\x00\x00\x3E\x42\x42\x3E\x02\x00", // "р" [0xD2]
	0x0441: "\x00\x00\x3E\x41\x01\x41\x3E\x00", // "с" [0xD3]
	0x0442: "\x00\x00\x7F\x08\x08\x08\x08\x00", // "т" [0xD4]
	0x0443: "\x00\x00\x42\x42\x7C\x40\x3C\x00", // "у" [0xD5]
	0x0436: "\x00\x00\x49\x49\x3E\x49\x49\x00", // "ж" [0xD6]
	0x0432: "\x00\x00\x3E\x42\x3E\x42\x3E\x00", // "в" [0xD7]
	0x044C: "\x00\x00\x02\x02\x3E\x42\x3E\x00", // "ь" [0xD8]
	0x044B: "\x00\x00\x21\x21\x27\x29\x27\x00", // "ы" [0xD9]
	0x0437: "\x00\x00\x3E\x40\x38\x40\x3E\x00", // "з" [0xDA]
	0x0448: "\x00\x00\x49\x49\x49\x49\x7F\x00", // "ш" [0xDB]
	0x044D: "\x00\x00\x3E\x40\x78\x40\x3E\x00", // "э" [0xDC]
	0x0449: "\x00\x00\x49\x49\x49\x7F\x40\x00", // "щ" [0xDD]
	0x0447: "\x00\x00\x42\x42\x7C\x40\x40\x00", // "ч" [0xDE]
	0x044A: "\x00\x00\x03\x02\x3E\x42\x3E\x00", // "ъ" [0xDF]
	0x042E: "\x31\x49\x49\x4F\x49\x49\x31\x00", // "Ю" [0xE0]
	0x0410: "\x1C\x22\x41\x7F\x41\x41\x41\x00", // "А" [0xE1]
	0x0411: "\x3F\x01\x01\x3F\x41\x41\x3F\x00", // "Б" [0xE2]
	0x0426: "\x21\x21\x21\x21\x21\x7F\x40\x00", // "Ц" [0xE3]
	0x0414: "\x3C\x22\x22\x22\x22\x7F\x41\x00", // "Д" [0xE4]
	0x0415: "\x7F\x01\x01\x1F\x01\x01\x7F\x00", // "Е" [0xE5]
	0x0424: "\x08\x3E\x49\x49\x49\x3E\x08\x00", // "Ф" [0xE6]
	0x0413: "\x7F\x01\x01\x01\x01\x01\x01\x00", // "Г" [0xE7]
	0x0425: "\x41\x22\x14\x08\x14\x22\x41\x00", // "Х" [0xE8]
	0x0418: "\x41\x61\x51\x49\x45\x43\x41\x00", // "И" [0xE9]
	0x0419: "\x14\x49\x61\x51\x49\x45\x43\x00", // "Й" [0xEA]
	0x041A: "\x21\x11\x09\x07\x09\x11\x21\x00", // "К" [0xEB]
	0x041B: "\x78\x44\x42\x42\x42\x42\x41\x00", // "Л" [0xEC]
	0x041C: "\x41\x63\x55\x49\x41\x41\x41\x00", // "М" [0xED]
	0x041D: "\x41\x41\x41\x7F\x41\x41\x41\x00", // "Н" [0xEE]
	0x041E: "\x3E\x41\x41\x41\x41\x41\x3E\x00", // "О" [0xEF]
	0x041F: "\x7F\x41\x41\x41\x41\x41\x41\x00", // "П" [0xF0]
	0x042F: "\x7E\x41\x41\x7E\x48\x44\x43\x00", // "Я" [0xF1]
	0x0420: "\x3F\x41\x41\x3F\x01\x01\x01\x00", // "Р" [0xF2]
	0x0421: "\x3E\x41\x01\x01\x01\x41\x3E\x00", // "С" [0xF3]
	0x0422: "\x7F\x08\x08\x08\x08\x08\x08\x00", // "Т" [0xF4]
	0x0423: "\x41\x41\x41\x7E\x40\x40\x3E\x00", // "У" [0xF5]
	0x0416: "\x49\x49\x49\x3E\x49\x49\x49\x00", // "Ж" [0xF6]
	0x0412: "\x3F\x41\x41\x3F\x41\x41\x3F\x00", // "В" [0xF7]
	0x042C: "\x01\x01\x01\x3F\x41\x41\x3F\x00", // "Ь" [0xF8]
	0x042B: "\x41\x41\x41\x4F\x51\x51\x4F\x00", // "Ы" [0xF9]
	0x0417: "\x3E\x41\x40\x38\x40\x41\x3E\x00", // "З" [0xFA]
	0x0428: "\x49\x49\x49\x49\x49\x49\x7F\x00", // "Ш" [0xFB]
	0x042D: "\x3E\x41\x40\x7C\x40\x41\x3E\x00", // "Э" [0xFC]
	0x0429: "\x49\x49\x49\x49\x49\x7F\x40\x00", // "Щ" [0xFD]
	0x0427: "\x41\x41\x41\x7E\x40\x40\x40\x00", // "Ч" [0xFE]
	0x042A: "\x03\x02\x02\x3E\x42\x42\x3E\x00", // "Ъ" [0xFF]
	0x20BD: "\x3F\x42\x42\x3F\x02\x1F\x02\x00", // "₽" [0x100]
	0x262D: "\x3E\x41\x4C\x46\x4B\x3C\x22\x41", // "☭" [0x101]
	0x263A: "\x00\x77\x66\x00\x00\x63\x3E\x00", // "☺" [0x102]
} // Font8
