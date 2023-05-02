// Font 8x8 source
// File: "font8src.go"

package font8src

// ----------------------------------------------------------------------------
var Font8src = []struct {
	cod int      // index (ASCII/KOI8-R code)
	chr string   // UTF-8 code
	val []string // [8]
}{
	{0x0020, " ", []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0021, "!", []string{
		"   #    ",
		"  ###   ",
		"  ###   ",
		"   #    ",
		"   #    ",
		"        ",
		"   #    ",
		"        "}},

	{0x0022, "\"", []string{
		"  # #   ",
		"  # #   ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0023, "#", []string{
		"  # #   ",
		"  # #   ",
		"####### ",
		"  # #   ",
		"####### ",
		"  # #   ",
		"  # #   ",
		"        "}},

	{0x0024, "$", []string{
		" #####  ",
		"#  #  # ",
		"#  #    ",
		" #####  ",
		"   #  # ",
		"#  #  # ",
		" #####  ",
		"        "}},

	{0x0025, "%", []string{
		"###   # ",
		"# #  #  ",
		"### #   ",
		"   #    ",
		"  # ### ",
		" #  # # ",
		"#   ### ",
		"        "}},

	{0x0026, "&", []string{
		"  ##    ",
		" #  #   ",
		"  ##    ",
		" ###    ",
		"#   # # ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0027, "'", []string{
		"   #    ",
		"   #    ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0028, "(", []string{
		"   ##   ",
		"  #     ",
		" #      ",
		" #      ",
		" #      ",
		"  #     ",
		"   ##   ",
		"        "}},

	{0x0029, ")", []string{
		"  ##    ",
		"    #   ",
		"     #  ",
		"     #  ",
		"     #  ",
		"    #   ",
		"  ##    ",
		"        "}},

	{0x002A, "*", []string{
		"        ",
		" #   #  ",
		"  # #   ",
		"### ### ",
		"  # #   ",
		" #   #  ",
		"        ",
		"        "}},

	{0x002B, "+", []string{
		"        ",
		"   #    ",
		"   #    ",
		" #####  ",
		"   #    ",
		"   #    ",
		"        ",
		"        "}},

	{0x002C, ",", []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"  ##    ",
		"   #    ",
		"  #     ",
		"        "}},

	{0x002D, "-", []string{
		"        ",
		"        ",
		"        ",
		" #####  ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x002E, ".", []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"   ##   ",
		"   ##   ",
		"        "}},

	{0x002F, "/", []string{
		"      # ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"#       ",
		"        "}},

	{0x0030, "0", []string{
		" #####  ",
		"#    ## ",
		"#   # # ",
		"#  #  # ",
		"# #   # ",
		"##    # ",
		" #####  ",
		"        "}},

	{0x0031, "1", []string{
		"   #    ",
		"  ##    ",
		" # #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		" #####  ",
		"        "}},

	{0x0032, "2", []string{
		" #####  ",
		"#     # ",
		"      # ",
		" #####  ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x0033, "3", []string{
		" #####  ",
		"#     # ",
		"      # ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0034, "4", []string{
		"    ##  ",
		"   # #  ",
		"  #  #  ",
		" #   #  ",
		"####### ",
		"     #  ",
		"     #  ",
		"        "}},

	{0x0035, "5", []string{
		"####### ",
		"#       ",
		"#       ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0036, "6", []string{
		" #####  ",
		"#     # ",
		"#       ",
		"######  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0037, "7", []string{
		"####### ",
		"#    #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		"  #     ",
		"  #     ",
		"        "}},

	{0x0038, "8", []string{
		" #####  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0039, "9", []string{
		" #####  ",
		"#     # ",
		"#     # ",
		" ###### ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x003A, ":", []string{
		"        ",
		"   ##   ",
		"   ##   ",
		"        ",
		"   ##   ",
		"   ##   ",
		"        ",
		"        "}},

	{0x003B, ";", []string{
		"        ",
		"   ##   ",
		"   ##   ",
		"        ",
		"   ##   ",
		"   ##   ",
		"   #    ",
		"        "}},

	{0x003C, "<", []string{
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"  #     ",
		"   #    ",
		"    #   ",
		"        "}},

	{0x003D, "=", []string{
		"        ",
		"        ",
		" #####  ",
		"        ",
		" #####  ",
		"        ",
		"        ",
		"        "}},

	{0x003E, ">", []string{
		"  #     ",
		"   #    ",
		"    #   ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		"        "}},

	{0x003F, "?", []string{
		" #####  ",
		"#     # ",
		"      # ",
		"    ##  ",
		"   #    ",
		"        ",
		"   #    ",
		"        "}},

	{0x0040, "@", []string{
		" #####  ",
		"#     # ",
		"#  #### ",
		"# #   # ",
		"#  ###  ",
		"#       ",
		" #####  ",
		"        "}},

	{0x0041, "A", []string{
		"  ###   ",
		" #   #  ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x0042, "B", []string{
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x0043, "C", []string{
		" #####  ",
		"#     # ",
		"#       ",
		"#       ",
		"#       ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0044, "D", []string{
		"#####   ",
		"#    #  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x0045, "E", []string{
		"####### ",
		"#       ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x0046, "F", []string{
		"####### ",
		"#       ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x0047, "G", []string{
		" #####  ",
		"#     # ",
		"#       ",
		"#  #### ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0048, "H", []string{
		"#     # ",
		"#     # ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x0049, "I", []string{
		"  ###   ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"  ###   ",
		"        "}},

	{0x004A, "J", []string{
		"    ### ",
		"     #  ",
		"     #  ",
		"     #  ",
		"     #  ",
		"#    #  ",
		" ####   ",
		"        "}},

	{0x004B, "K", []string{
		"#    #  ",
		"#   #   ",
		"#  #    ",
		"###     ",
		"#  #    ",
		"#   #   ",
		"#    #  ",
		"        "}},

	{0x004C, "L", []string{
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x004D, "M", []string{
		"#     # ",
		"##   ## ",
		"# # # # ",
		"#  #  # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x004E, "N", []string{
		"#     # ",
		"##    # ",
		"# #   # ",
		"#  #  # ",
		"#   # # ",
		"#    ## ",
		"#     # ",
		"        "}},

	{0x004F, "O", []string{
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0050, "P", []string{
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x0051, "Q", []string{
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#   # # ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0052, "R", []string{
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#   #   ",
		"#    #  ",
		"#     # ",
		"        "}},

	{0x0053, "S", []string{
		" #####  ",
		"#     # ",
		"#       ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0054, "T", []string{
		"####### ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x0055, "U", []string{
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0056, "V", []string{
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"        "}},

	{0x0057, "W", []string{
		"#     # ",
		"#     # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		" ## ##  ",
		"        "}},

	{0x0058, "X", []string{
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"  # #   ",
		" #   #  ",
		"#     # ",
		"        "}},

	{0x0059, "Y", []string{
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x005A, "Z", []string{
		"####### ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"####### ",
		"        "}},

	{0x005B, "[", []string{
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #####  ",
		"        "}},

	{0x005C, "\\", []string{
		"#       ",
		" #      ",
		"  #     ",
		"   #    ",
		"    #   ",
		"     #  ",
		"      # ",
		"        "}},

	{0x005D, "]", []string{
		" #####  ",
		"     #  ",
		"     #  ",
		"     #  ",
		"     #  ",
		"     #  ",
		" #####  ",
		"        "}},

	{0x005E, "^", []string{
		"   #    ",
		"  # #   ",
		" #   #  ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x005F, "_", []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"####### ",
		"        "}},

	{0x0060, "`", []string{
		"  ##    ",
		"    ##  ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0061, "a", []string{
		"        ",
		"        ",
		" #####  ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0062, "b", []string{
		"#       ",
		"#       ",
		"######  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"######  ",
		"         "}},

	{0x0063, "c", []string{
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"#       ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0064, "d", []string{
		"     #  ",
		"     #  ",
		" #####  ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0065, "e", []string{
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"####### ",
		"#       ",
		" #####  ",
		"        "}},

	{0x0066, "f", []string{
		"   ##   ",
		"  #  #  ",
		"  #     ",
		" ###    ",
		"  #     ",
		"  #     ",
		"  #     ",
		"        "}},

	{0x0067, "g", []string{
		"        ",
		"  ####  ",
		" #    # ",
		" #    # ",
		"  ##### ",
		"      # ",
		"  ####  ",
		"        "}},

	{0x0068, "h", []string{
		" #      ",
		" #      ",
		" # ###  ",
		" ##   # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"        "}},

	{0x0069, "i", []string{
		"   #    ",
		"        ",
		"  ##    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"  ###   ",
		"        "}},

	{0x006A, "j", []string{
		"    #   ",
		"        ",
		"   ##   ",
		"    #   ",
		"    #   ",
		" #  #   ",
		"  ##    ",
		"        "}},

	{0x006B, "k", []string{
		" #      ",
		" #      ",
		" #   #  ",
		" #  #   ",
		" ###    ",
		" #  #   ",
		" #   #  ",
		"        "}},

	{0x006C, "l", []string{
		"  ##    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"  ###   ",
		"        "}},

	{0x006D, "m", []string{
		"        ",
		"        ",
		"### ##  ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"        "}},

	{0x006E, "n", []string{
		"        ",
		"        ",
		" # ###  ",
		" ##   # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"        "}},

	{0x006F, "o", []string{
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0070, "p", []string{
		"        ",
		"        ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #      ",
		"        "}},

	{0x0071, "q", []string{
		"        ",
		"        ",
		"  ##### ",
		" #    # ",
		" #    # ",
		"  ##### ",
		"      # ",
		"        "}},

	{0x0072, "r", []string{
		"        ",
		"        ",
		" # ###  ",
		" ##   # ",
		" #      ",
		" #      ",
		" #      ",
		"        "}},

	{0x0073, "s", []string{
		"        ",
		"        ",
		"  ####  ",
		" #      ",
		"  ####  ",
		"      # ",
		" #####  ",
		"        "}},

	{0x0074, "t", []string{
		"  #     ",
		"  #     ",
		" ###    ",
		"  #     ",
		"  #     ",
		"  #  #  ",
		"   ##   ",
		"        "}},

	{0x0075, "u", []string{
		"        ",
		"        ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #   ## ",
		"  ### # ",
		"        "}},

	{0x0076, "v", []string{
		"        ",
		"        ",
		"#     # ",
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"        "}},

	{0x0077, "w", []string{
		"        ",
		"        ",
		"#     # ",
		"#     # ",
		"#  #  # ",
		"#  #  # ",
		" ## ##  ",
		"       "}},

	{0x0078, "x", []string{
		"        ",
		"        ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"  # #   ",
		" #   #  ",
		"        "}},

	{0x0079, "y", []string{
		"        ",
		"        ",
		" #    # ",
		" #    # ",
		"  ##### ",
		"      # ",
		"  ####  ",
		"        "}},

	{0x007A, "z", []string{
		"        ",
		"        ",
		" #####  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #####  ",
		"        "}},

	{0x007B, "{", []string{
		"   ###  ",
		"  #     ",
		"  #     ",
		" ##     ",
		"  #     ",
		"  #     ",
		"   ###  ",
		"        "}},

	{0x007C, "|", []string{
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x007D, "}", []string{
		" ###    ",
		"    #   ",
		"    #   ",
		"    ##  ",
		"    #   ",
		"    #   ",
		" ###    ",
		"        "}},

	{0x007E, "~", []string{
		"        ",
		"        ",
		" ##     ",
		"#  #  # ",
		"    ##  ",
		"        ",
		"        ",
		"        "}},

	{0x007F, "\u007F", []string{ // [DEL]
		"# # # # ",
		" # # # #",
		"# # # # ",
		" # # # #",
		"# # # # ",
		" # # # #",
		"# # # # ",
		"        "}},

	{0x0080, "\u2500", []string{ // "─"
		"        ",
		"        ",
		"        ",
		"########",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0081, "\u2502", []string{ // "│"
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0082, "\u250C", []string{ // "┌"
		"        ",
		"        ",
		"        ",
		"   #####",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0083, "\u2510", []string{ // "┐"
		"        ",
		"        ",
		"        ",
		"####    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0084, "\u2514", []string{ // "└"
		"   #    ",
		"   #    ",
		"   #    ",
		"   #####",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0085, "\u2518", []string{ // "┘"
		"   #    ",
		"   #    ",
		"   #    ",
		"####    ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0086, "\u251C", []string{ // "├"
		"   #    ",
		"   #    ",
		"   #    ",
		"   #####",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0087, "\u2524", []string{ // "┤"
		"   #    ",
		"   #    ",
		"   #    ",
		"####    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0088, "\u252C", []string{ // "┬"
		"        ",
		"        ",
		"        ",
		"########",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0089, "\u2534", []string{ // "┴"
		"   #    ",
		"   #    ",
		"   #    ",
		"########",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x008A, "\u253C", []string{ // "┼"
		"   #    ",
		"   #    ",
		"   #    ",
		"########",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x008B, "\u2580", []string{ // "▀"
		"########",
		"########",
		"########",
		"########",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x008C, "\u2584", []string{ // "▄"
		"        ",
		"        ",
		"        ",
		"        ",
		"########",
		"########",
		"########",
		"########"}},

	{0x008D, "\u2588", []string{ // "█"
		"########",
		"########",
		"########",
		"########",
		"########",
		"########",
		"########",
		"########"}},

	{0x008E, "\u258C", []string{ // "▌"
		"####    ",
		"####    ",
		"####    ",
		"####    ",
		"####    ",
		"####    ",
		"####    ",
		"####    "}},

	{0x008F, "\u2590", []string{ // "▐"
		"    ####",
		"    ####",
		"    ####",
		"    ####",
		"    ####",
		"    ####",
		"    ####",
		"    ####"}},

	{0x0090, "\u2591", []string{ // "░"
		"#  #  # ",
		"        ",
		" #  #  #",
		"        ",
		"#  #  # ",
		"        ",
		" #  #  #",
		"        "}},

	{0x0091, "\u2592", []string{ // "▒"
		"#  #  # ",
		" #  #  #",
		"  #  #  ",
		"#  #  # ",
		" #  #  #",
		"  #  #  ",
		"#  #  # ",
		" #  #  #"}},

	{0x0092, "\u2593", []string{ // "▓"
		"# # # # ",
		" # # # #",
		"# # # # ",
		" # # # #",
		"# # # # ",
		" # # # #",
		"# # # # ",
		" # # # #"}},

	{0x0093, "\u2320", []string{ // "⌠"
		"    ##  ",
		"   #  # ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x0094, "\u25A0", []string{ // "■"
		"        ",
		"####### ",
		"####### ",
		"####### ",
		"####### ",
		"####### ",
		"        ",
		"        "}},

	{0x0095, "\u2219", []string{ // "∙"
		"        ",
		"        ",
		"   ##   ",
		"  ####  ",
		"   ##   ",
		"        ",
		"        ",
		"        "}},

	{0x0096, "\u221A", []string{ // "√"
		"    ### ",
		"    #   ",
		"#   #   ",
		" #  #   ",
		"  # #   ",
		"   ##    ",
		"    #   ",
		"        "}},

	{0x0097, "\u2248", []string{ // "≈"
		"        ",
		"  ##  # ",
		" #  ##  ",
		"        ",
		"  ##  # ",
		" #  ##  ",
		"        ",
		"        "}},

	{0x0098, "\u2264", []string{ // "≤"
		"        ",
		"   #    ",
		"  #     ",
		" #      ",
		"  #     ",
		"   #    ",
		"        ",
		" ###    "}},

	{0x0099, "\u2265", []string{ // "≥"
		"        ",
		" #      ",
		"  #     ",
		"   #    ",
		"  #     ",
		" #      ",
		"        ",
		" ###    "}},

	{0x009A, "\u00A0", []string{ // " "
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x009B, "\u2321", []string{ // "⌡"
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"#  #    ",
		" ##     "}},

	{0x009C, "\u00B0", []string{ // "°"
		"   ##   ",
		"  #  #  ",
		"  #  #  ",
		"   ##   ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x009D, "\u00B2", []string{ // "²"
		"   ###  ",
		"  #   # ",
		"    ##  ",
		"   #    ",
		"  ##### ",
		"        ",
		"        ",
		"        "}},

	{0x009E, "\u00B7", []string{ // "·"
		"        ",
		"        ",
		"        ",
		"   #    ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x009F, "\u00F7", []string{ // "÷"
		"        ",
		"   #    ",
		"        ",
		" #####  ",
		"        ",
		"   #    ",
		"        ",
		"        "}},

	{0x00A0, "\u2550", []string{ // "═"
		"        ",
		"        ",
		"########",
		"        ",
		"########",
		"        ",
		"        ",
		"        "}},

	{0x00A1, "\u2551", []string{ // "║"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00A2, "\u2552", []string{ // "╒"
		"        ",
		"        ",
		"   #####",
		"   #    ",
		"   #####",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00A3, "\u0451", []string{ // "ё"
		"  # #   ",
		"        ",
		" #####  ",
		"#     # ",
		"####### ",
		"#       ",
		" #####  ",
		"        "}},

	{0x00A4, "\u2553", []string{ // "╓"
		"        ",
		"        ",
		"        ",
		"  ######",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00A5, "\u2554", []string{ // "╔"
		"        ",
		"        ",
		"  ######",
		"  #     ",
		"  # ####",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00A6, "\u2555", []string{ // "╕"
		"        ",
		"        ",
		"####    ",
		"   #    ",
		"####    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00A7, "\u2556", []string{ // "╖"
		"        ",
		"        ",
		"        ",
		"#####   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00A8, "\u2557", []string{ // "╗"
		"        ",
		"        ",
		"#####   ",
		"    #   ",
		"### #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00A9, "\u2558", []string{ // "╘"
		"   #    ",
		"   #    ",
		"   #####",
		"   #    ",
		"   #####",
		"        ",
		"        ",
		"        "}},

	{0x00AA, "\u2559", []string{ // "╙"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  ######",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x00AB, "\u255A", []string{ // "╚"
		"  # #   ",
		"  # #   ",
		"  # ####",
		"  #     ",
		"  ######",
		"        ",
		"        ",
		"        "}},

	{0x00AC, "\u255B", []string{ // "╛"
		"   #    ",
		"   #    ",
		"####    ",
		"   #    ",
		"####    ",
		"        ",
		"        ",
		"        "}},

	{0x00AD, "\u255C", []string{ // "╜"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"#####   ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x00AE, "\u255D", []string{ // "╝"
		"  # #   ",
		"  # #   ",
		"### #   ",
		"    #   ",
		"#####   ",
		"        ",
		"        ",
		"        "}},

	{0x00AF, "\u255E", []string{ // "╞"
		"   #    ",
		"   #    ",
		"   #####",
		"   #    ",
		"   #####",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00B0, "\u255F", []string{ // "╟"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # ####",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B1, "\u2560", []string{ // "╠"
		"  # #   ",
		"  # #   ",
		"  # ####",
		"  #     ",
		"  # ####",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B2, "\u2561", []string{ // "╡"
		"   #    ",
		"   #    ",
		"####    ",
		"   #    ",
		"####    ",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00B3, "\u0401", []string{ // "Ё"
		"  # #  ",
		"####### ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x00B4, "\u2562", []string{ // "╢"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"### #   ",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B5, "\u2563", []string{ // "╣"
		"  # #   ",
		"  # #   ",
		"### #   ",
		"    #   ",
		"### #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B6, "\u2564", []string{ // "╤"
		"        ",
		"        ",
		"########",
		"        ",
		"########",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00B7, "\u2565", []string{ // "╥"
		"        ",
		"        ",
		"        ",
		"########",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B8, "\u2566", []string{ // "╦"
		"        ",
		"        ",
		"########",
		"        ",
		"### ####",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00B9, "\u2567", []string{ // "╧"
		"   #    ",
		"   #    ",
		"########",
		"        ",
		"########",
		"        ",
		"        ",
		"        "}},

	{0x00BA, "\u2568", []string{ // "╨"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"########",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x00BB, "\u2569", []string{ // "╩"
		"  # #   ",
		"  # #   ",
		"### ####",
		"        ",
		"########",
		"        ",
		"        ",
		"        "}},

	{0x00BC, "\u256A", []string{ // "╪"
		"   #    ",
		"   #    ",
		"########",
		"        ",
		"########",
		"   #    ",
		"   #    ",
		"   #    "}},

	{0x00BD, "\u256B", []string{ // "╫"
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"########",
		"  # #   ",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00BE, "\u256C", []string{ // "╬"
		"  # #   ",
		"  # #   ",
		"### ####",
		"        ",
		"### ####",
		"  # #   ",
		"  # #   ",
		"  # #   "}},

	{0x00BF, "\u00A9", []string{ // "©"
		" ###### ",
		"#      #",
		"# #### #",
		"# #    #",
		"# #### #",
		"#      #",
		" ###### ",
		"        "}},

	{0x00C0, "\u044E", []string{ // "ю"
		"        ",
		"        ",
		"#   ##  ",
		"#  #  # ",
		"####  # ",
		"#  #  # ",
		"#   ##  ",
		"        "}},

	{0x00C1, "\u0430", []string{ // "a"
		"        ",
		"        ",
		" #####  ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x00C2, "\u0431", []string{ // "б"
		" #####  ",
		"#       ",
		"# ####  ",
		"##    # ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x00C3, "\u0446", []string{ // "ц"
		"        ",
		"        ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		"####### ",
		"      # ",
		"        "}},

	{0x00C4, "\u0434", []string{ // "д"
		"        ",
		"        ",
		"  ####  ",
		" #   #  ",
		" #   #  ",
		"####### ",
		"#     # ",
		"        "}},

	{0x00C5, "\u0435", []string{ // "е"
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"####### ",
		"#       ",
		" #####  ",
		"        "}},

	{0x00C6, "\u0444", []string{ // "ф"
		"        ",
		"   #    ",
		" #####  ",
		"#  #  # ",
		" #####  ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x00C7, "\u0433", []string{ // "г"
		"        ",
		"        ",
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		"        "}},

	{0x0C8, "\u0445", []string{ // "х"
		"        ",
		"        ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"  # #   ",
		" #   #  ",
		"        "}},

	{0x00C9, "\u0438", []string{ // "и"
		"        ",
		"        ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #   ## ",
		"  ### # ",
		"        "}},

	{0x00CA, "\u0439", []string{ // "й"
		"  #  #  ",
		"   ##   ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #   ## ",
		"  ### # ",
		"        "}},

	{0x00CB, "\u043A", []string{ // "к"
		"        ",
		"        ",
		" #  ##  ",
		" # #    ",
		" ###    ",
		" #  #   ",
		" #   #  ",
		"        "}},

	{0x00CC, "\u043B", []string{ // "л"
		"        ",
		"        ",
		"   #### ",
		"  #   # ",
		"  #   # ",
		"  #   # ",
		" #    # ",
		"        "}},

	{0x00CD, "\u043C", []string{ // "м"
		"        ",
		"        ",
		"#     # ",
		"##   ## ",
		"# # # # ",
		"#  #  # ",
		"#     # ",
		"        "}},

	{0x00CE, "\u043D", []string{ // "н"
		"        ",
		"        ",
		"#     # ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00CF, "\u043E", []string{ // "о"
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00D0, "\u043F", []string{ // "п"
		"        ",
		"        ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00D1, "\u044F", []string{ // "я"
		"        ",
		"        ",
		" ###### ",
		"#     # ",
		" ###### ",
		"  #   # ",
		"##    # ",
		"        "}},

	{0x00D2, "\u0440", []string{ // "р"
		"        ",
		"        ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #      ",
		"        "}},

	{0x00D3, "\u0441", []string{ // "с"
		"        ",
		"        ",
		" #####  ",
		"#     # ",
		"#       ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00D4, "\u0442", []string{ // "т"
		"        ",
		"        ",
		"####### ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x00D5, "\u0443", []string{ // "у"
		"        ",
		"        ",
		" #    # ",
		" #    # ",
		"  ##### ",
		"      # ",
		"  ####  ",
		"        "}},

	{0x00D6, "\u0436", []string{ // "ж"
		"        ",
		"        ",
		"#  #  # ",
		"#  #  # ",
		" #####  ",
		"#  #  # ",
		"#  #  # ",
		"        "}},

	{0x00D7, "\u0432", []string{ // "в"
		"        ",
		"        ",
		" #####  ",
		" #    # ",
		" #####  ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x00D8, "\u044C", []string{ // "ь"
		"        ",
		"        ",
		" #      ",
		" #      ",
		" #####  ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x00D9, "\u044B", []string{ // "ы"
		"       ",
		"       ",
		"#    # ",
		"#    # ",
		"###  # ",
		"#  # # ",
		"###  # ",
		"       "}},

	{0x00DA, "\u0437", []string{ // "з"
		"        ",
		"        ",
		" #####  ",
		"      # ",
		"   ###  ",
		"      # ",
		" #####  ",
		"        "}},

	{0x00DB, "\u0448", []string{ // "ш"
		"        ",
		"        ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"####### ",
		"        "}},

	{0x00DC, "\u044D", []string{ // "э"
		"        ",
		"        ",
		" #####  ",
		"      # ",
		"   #### ",
		"      # ",
		" #####  ",
		"        "}},

	{0x00DD, "\u0449", []string{ // "щ"
		"        ",
		"        ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"####### ",
		"      # ",
		"        "}},

	{0x00DE, "\u0447", []string{ // "ч"
		"        ",
		"        ",
		" #    # ",
		" #    # ",
		"  ##### ",
		"      # ",
		"      # ",
		"        "}},

	{0x00DF, "\u044A", []string{ // "ъ"
		"        ",
		"        ",
		"##      ",
		" #      ",
		" #####  ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x00E0, "\u042E", []string{ // "Ю"
		"#   ##  ",
		"#  #  # ",
		"#  #  # ",
		"####  # ",
		"#  #  # ",
		"#  #  # ",
		"#   ##  ",
		"        "}},

	{0x00E1, "\u0410", []string{ // "А"
		"  ###   ",
		" #   #  ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00E2, "\u0411", []string{ // "Б"
		"######  ",
		"#       ",
		"#       ",
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x00E3, "\u0426", []string{ // "Ц"
		"#    #  ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		"#    #  ",
		"####### ",
		"      # ",
		"        "}},

	{0x00E4, "\u0414", []string{ // "Д"
		"  ####  ",
		" #   #  ",
		" #   #  ",
		" #   #  ",
		" #   #  ",
		"####### ",
		"#     # ",
		"        "}},

	{0x00E5, "\u0415", []string{ // "Е"
		"####### ",
		"#       ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x00E6, "\u0424", []string{ // "Ф"
		"   #    ",
		" #####  ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		" #####  ",
		"   #    ",
		"        "}},

	{0x00E7, "\u0413", []string{ // "Г"
		"####### ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x00E8, "\u0425", []string{ // "Х"
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"  # #   ",
		" #   #  ",
		"#     # ",
		"        "}},

	{0x00E9, "\u0418", []string{ // "И"
		"#     # ",
		"#    ## ",
		"#   # # ",
		"#  #  # ",
		"# #   # ",
		"##    # ",
		"#     # ",
		"        "}},

	{0x00EA, "\u0419", []string{ // "Й"
		"  # #   ",
		"#  #  # ",
		"#    ## ",
		"#   # # ",
		"#  #  # ",
		"# #   # ",
		"##    # ",
		"        "}},

	{0x00EB, "\u041A", []string{ // "К"
		"#    #  ",
		"#   #   ",
		"#  #    ",
		"###     ",
		"#  #    ",
		"#   #   ",
		"#    #  ",
		"        "}},

	{0x00EC, "\u041B", []string{ // "Л"
		"   #### ",
		"  #   # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"#     # ",
		"        "}},

	{0x00ED, "\u041C", []string{ // "М"
		"#     # ",
		"##   ## ",
		"# # # # ",
		"#  #  # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00EE, "\u041D", []string{ // "Н"
		"#     # ",
		"#     # ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00EF, "\u041E", []string{ // "О"
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00F0, "\u041F", []string{ // "П"
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x00F1, "\u042F", []string{ // "Я"
		" ###### ",
		"#     # ",
		"#     # ",
		" ###### ",
		"   #  # ",
		"  #   # ",
		"##    # ",
		"        "}},

	{0x00F2, "\u0420", []string{ // "Р"
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x00F3, "\u0421", []string{ // "С"
		" #####  ",
		"#     # ",
		"#       ",
		"#       ",
		"#       ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00F4, "\u0422", []string{ // "Т"
		"####### ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x00F5, "\u0423", []string{ // "У"
		"#     # ",
		"#     # ",
		"#     # ",
		" ###### ",
		"      # ",
		"      # ",
		" #####  ",
		"        "}},

	{0x00F6, "\u0416", []string{ // "Ж"
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		" #####  ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"        "}},

	{0x00F7, "\u0412", []string{ // "В"
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x00F8, "\u042C", []string{ // "Ь"
		"#       ",
		"#       ",
		"#       ",
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x00F9, "\u042B", []string{ // "Ы"
		"#     # ",
		"#     # ",
		"#     # ",
		"####  # ",
		"#   # # ",
		"#   # # ",
		"####  # ",
		"        "}},

	{0x00FA, "\u0417", []string{ // "З"
		" #####  ",
		"#     # ",
		"      # ",
		"   ###  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00FB, "\u0428", []string{ // "Ш"
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"####### ",
		"        "}},

	{0x00FC, "\u042D", []string{ // "Э"
		" #####  ",
		"#     # ",
		"      # ",
		"  ##### ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x00FD, "\u0429", []string{ // "Щ"
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"####### ",
		"      # ",
		"        "}},

	{0x00FE, "\u0427", []string{ // "Ч"
		"#     # ",
		"#     # ",
		"#     # ",
		" ###### ",
		"      # ",
		"      # ",
		"      # ",
		"        "}},

	{0x00FF, "\u042A", []string{ // "Ъ"
		"##      ",
		" #      ",
		" #      ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x0100, "\u20BD", []string{ // "₽"
		"######  ",
		" #    # ",
		" #    # ",
		"######  ",
		" #      ",
		"#####   ",
		" #      ",
		"        "}},

	{0x0101, "\u262D", []string{ // "☭"
		" #####  ",
		"#     # ",
		"  ##  # ",
		" ##   # ",
		"## #  # ",
		"  ####  ",
		" #   #  ",
		"#     # "}},

	{0x0102, "\u263A", []string{ // "☺"
		"        ",
		"### ### ",
		" ##  ## ",
		"        ",
		"        ",
		"##   ## ",
		" #####  ",
		"        "}},
} // Font8src
// ----------------------------------------------------------------------------

/*** end of "font8src.go" file ***/