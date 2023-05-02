var Font = []struct {
	cod int      // index (ASCII/KOI8-R code)
	chr string   // UTF-8 code
	val []string // [8]
}{
	{0x0020, "\u0020", []string{ // " "
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0021, "\u0021", []string{ // "!"
		"  ###   ",
		"  ###   ",
		"  ###   ",
		"   #    ",
		"        ",
		"  ###   ",
		"  ###   ",
		"        "}},

	{0x0022, "\u0022", []string{ // "\""
		"### ### ",
		"### ### ",
		" #   #  ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0023, "\u0023", []string{ // "#"
		"  # #   ",
		"  # #   ",
		"####### ",
		"  # #   ",
		"####### ",
		"  # #   ",
		"  # #   ",
		"        "}},

	{0x0024, "\u0024", []string{ // "$"
		" #####  ",
		"#  #  # ",
		"#  #    ",
		" #####  ",
		"   #  # ",
		"#  #  # ",
		" #####  ",
		"        "}},

	{0x0025, "\u0025", []string{ // "%"
		"###   # ",
		"# #  #  ",
		"### #   ",
		"   #    ",
		"  # ### ",
		" #  # # ",
		"#   ### ",
		"        "}},

	{0x0026, "\u0026", []string{ // "&"
		"  ##    ",
		" #  #   ",
		"  ##    ",
		" ###    ",
		"#   # # ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0027, "\u0027", []string{ // "'"
		"  ###   ",
		"  ###   ",
		"   #    ",
		"  #     ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0028, "\u0028", []string{ // "("
		"   ##   ",
		"  #     ",
		" #      ",
		" #      ",
		" #      ",
		"  #     ",
		"   ##   ",
		"        "}},

	{0x0029, "\u0029", []string{ // ")"
		"  ##    ",
		"    #   ",
		"     #  ",
		"     #  ",
		"     #  ",
		"    #   ",
		"  ##    ",
		"        "}},

	{0x002A, "\u002A", []string{ // "*"
		"        ",
		" #   #  ",
		"  # #   ",
		"### ### ",
		"  # #   ",
		" #   #  ",
		"        ",
		"        "}},

	{0x002B, "\u002B", []string{ // "+"
		"        ",
		"   #    ",
		"   #    ",
		" #####  ",
		"   #    ",
		"   #    ",
		"        ",
		"        "}},

	{0x002C, "\u002C", []string{ // ","
		"        ",
		"        ",
		"        ",
		"  ###   ",
		"  ###   ",
		"   #    ",
		"  #     ",
		"        "}},

	{0x002D, "\u002D", []string{ // "-"
		"        ",
		"        ",
		"        ",
		" #####  ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x002E, "\u002E", []string{ // "."
		"        ",
		"        ",
		"        ",
		"        ",
		"  ###   ",
		"  ###   ",
		"  ###   ",
		"        "}},

	{0x002F, "\u002F", []string{ // "/"
		"      # ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"#       ",
		"        "}},

	{0x0030, "\u0030", []string{ // "0"
		"  ###   ",
		" #   #  ",
		"# #   # ",
		"#  #  # ",
		"#   # # ",
		" #   #  ",
		"  ###   ",
		"        "}},

	{0x0031, "\u0031", []string{ // "1"
		"   #    ",
		"  ##    ",
		" # #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		" #####  ",
		"        "}},

	{0x0032, "\u0032", []string{ // "2"
		" #####  ",
		"#     # ",
		"      # ",
		" #####  ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x0033, "\u0033", []string{ // "3"
		" #####  ",
		"#     # ",
		"      # ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0034, "\u0034", []string{ // "4"
		"#       ",
		"#    #  ",
		"#    #  ",
		"####### ",
		"     #  ",
		"     #  ",
		"     #  ",
		"        "}},

	{0x0035, "\u0035", []string{ // "5"
		"####### ",
		"#       ",
		"#       ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0036, "\u0036", []string{ // "6"
		" #####  ",
		"#     # ",
		"#       ",
		"######  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0037, "\u0037", []string{ // "7"
		"####### ",
		"#    #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		"  #     ",
		"  #     ",
		"        "}},

	{0x0038, "\u0038", []string{ // "8"
		" #####  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0039, "\u0039", []string{ // "9"
		" #####  ",
		"#     # ",
		"#     # ",
		" ###### ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x003A, "\u003A", []string{ // ":"
		"   #    ",
		"  # #   ",
		"   #    ",
		"        ",
		"   #    ",
		"  # #   ",
		"   #    ",
		"        "}},

	{0x003B, "\u003B", []string{ // ";"
		"  ###   ",
		"  ###   ",
		"        ",
		"  ###   ",
		"  ###   ",
		"   #    ",
		"  #     ",
		"        "}},

	{0x003C, "\u003C", []string{ // "<"
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"  #     ",
		"   #    ",
		"    #   ",
		"        "}},

	{0x003D, "\u003D", []string{ // "="
		"        ",
		"        ",
		" #####  ",
		"        ",
		" #####  ",
		"        ",
		"        ",
		"        "}},

	{0x003E, "\u003E", []string{ // ">"
		"  #     ",
		"   #    ",
		"    #   ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		"        "}},

	{0x003F, "\u003F", []string{ // "?"
		" #####  ",
		"#     # ",
		"      # ",
		"    ##  ",
		"   #    ",
		"        ",
		"   #    ",
		"        "}},

	{0x0040, "\u0040", []string{ // "@"
		" #####  ",
		"#     # ",
		"# ### # ",
		"# # # # ",
		"# ####  ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0041, "\u0041", []string{ // "A"
		"   #    ",
		"  # #   ",
		" #   #  ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x0042, "\u0042", []string{ // "B"
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x0043, "\u0043", []string{ // "C"
		" #####  ",
		"#     # ",
		"#       ",
		"#       ",
		"#       ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0044, "\u0044", []string{ // "D"
		"######  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"######  ",
		"        "}},

	{0x0045, "\u0045", []string{ // "E"
		"####### ",
		"#       ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x0046, "\u0046", []string{ // "F"
		"####### ",
		"#       ",
		"#       ",
		"#####   ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x0047, "\u0047", []string{ // "G"
		" #####  ",
		"#     # ",
		"#       ",
		"#  #### ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0048, "\u0048", []string{ // "H"
		"#     # ",
		"#     # ",
		"#     # ",
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x0049, "\u0049", []string{ // "I"
		"  ###   ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"  ###   ",
		"        "}},

	{0x004A, "\u004A", []string{ // "J"
		"      # ",
		"      # ",
		"      # ",
		"      # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x004B, "\u004B", []string{ // "K"
		"#    #  ",
		"#   #   ",
		"#  #    ",
		"###     ",
		"#  #    ",
		"#   #   ",
		"#    #  ",
		"        "}},

	{0x004C, "\u004C", []string{ // "L"
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"#       ",
		"####### ",
		"        "}},

	{0x004D, "\u004D", []string{ // "M"
		"#     # ",
		"##   ## ",
		"# # # # ",
		"#  #  # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"        "}},

	{0x004E, "\u004E", []string{ // "N"
		"#     # ",
		"##    # ",
		"# #   # ",
		"#  #  # ",
		"#   # # ",
		"#    ## ",
		"#     # ",
		"        "}},

	{0x004F, "\u004F", []string{ // "O"
		"####### ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"####### ",
		"        "}},

	{0x0050, "\u0050", []string{ // "P"
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#       ",
		"#       ",
		"#       ",
		"        "}},

	{0x0051, "\u0051", []string{ // "Q"
		" #####  ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#   # # ",
		"#    #  ",
		" #### # ",
		"        "}},

	{0x0052, "\u0052", []string{ // "R"
		"######  ",
		"#     # ",
		"#     # ",
		"######  ",
		"#   #   ",
		"#    #  ",
		"#     # ",
		"        "}},

	{0x0053, "\u0053", []string{ // "S"
		" #####  ",
		"#     # ",
		"#       ",
		" #####  ",
		"      # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0054, "\u0054", []string{ // "T"
		"####### ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x0055, "\u0055", []string{ // "U"
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #####  ",
		"        "}},

	{0x0056, "\u0056", []string{ // "V"
		"#     # ",
		"#     # ",
		"#     # ",
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"        "}},

	{0x0057, "\u0057", []string{ // "W"
		"#     # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		"#  #  # ",
		" ## ##  ",
		"        "}},

	{0x0058, "\u0058", []string{ // "X"
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"  # #   ",
		" #   #  ",
		"#     # ",
		"        "}},

	{0x0059, "\u0059", []string{ // "Y"
		"#     # ",
		" #   #  ",
		"  # #   ",
		"   #    ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x005A, "\u005A", []string{ // "Z"
		"####### ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #      ",
		"####### ",
		"        "}},

	{0x005B, "\u005B", []string{ // "["
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #####  ",
		"        "}},

	{0x005C, "\u005C", []string{ // "\\"
		"#       ",
		" #      ",
		"  #     ",
		"   #    ",
		"    #   ",
		"     #  ",
		"      # ",
		"        "}},

	{0x005D, "\u005D", []string{ // "]"
		" #####  ",
		"     #  ",
		"     #  ",
		"     #  ",
		"     #  ",
		"     #  ",
		" #####  ",
		"        "}},

	{0x005E, "\u005E", []string{ // "^"
		"   #    ",
		"  # #   ",
		" #   #  ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x005F, "\u005F", []string{ // "_"
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"####### ",
		"        "}},

	{0x0060, "\u0060", []string{ // "`"
		"  ###   ",
		"  ###   ",
		"   #    ",
		"    #   ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x0061, "\u0061", []string{ // "a"
		"        ",
		"   ##   ",
		"  #  #  ",
		" #    # ",
		" ###### ",
		" #    # ",
		" #    # ",
		"        "}},

	{0x0062, "\u0062", []string{ // "b"
		"        ",
		" #####  ",
		" #    # ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x0063, "\u0063", []string{ // "c"
		"        ",
		"  ####  ",
		" #    # ",
		" #      ",
		" #      ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x0064, "\u0064", []string{ // "d"
		"        ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #####  ",
		"        "}},

	{0x0065, "\u0065", []string{ // "e"
		"        ",
		" ###### ",
		" #      ",
		" #####  ",
		" #      ",
		" #      ",
		" ###### ",
		"        "}},

	{0x0066, "\u0066", []string{ // "f"
		"        ",
		" ###### ",
		" #      ",
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		"        "}},

	{0x0067, "\u0067", []string{ // "g"
		"        ",
		"  ####  ",
		" #    # ",
		" #      ",
		" #  ### ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x0068, "\u0068", []string{ // "h"
		"        ",
		" #    # ",
		" #    # ",
		" ###### ",
		" #    # ",
		" #    # ",
		" #    # ",
		"        "}},

	{0x0069, "\u0069", []string{ // "i"
		"        ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"        "}},

	{0x006A, "\u006A", []string{ // "j"
		"        ",
		"      # ",
		"      # ",
		"      # ",
		"      # ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x006B, "\u006B", []string{ // "k"
		"        ",
		" #    # ",
		" #   #  ",
		" ####   ",
		" #  #   ",
		" #   #  ",
		" #    # ",
		"        "}},

	{0x006C, "\u006C", []string{ // "l"
		"        ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" ###### ",
		"        "}},

	{0x006D, "\u006D", []string{ // "m"
		"        ",
		" #    # ",
		" ##  ## ",
		" # ## # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"        "}},

	{0x006E, "\u006E", []string{ // "n"
		"        ",
		" #    # ",
		" ##   # ",
		" # #  # ",
		" #  # # ",
		" #   ## ",
		" #    # ",
		"        "}},

	{0x006F, "\u006F", []string{ // "o"
		"        ",
		"  ####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x0070, "\u0070", []string{ // "p"
		"        ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #      ",
		" #      ",
		"        "}},

	{0x0071, "\u0071", []string{ // "q"
		"        ",
		"  ####  ",
		" #    # ",
		" #    # ",
		" #  # # ",
		" #   #  ",
		"  ### # ",
		"        "}},

	{0x0072, "\u0072", []string{ // "r"
		"        ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #   #  ",
		" #    # ",
		"        "}},

	{0x0073, "\u0073", []string{ // "s"
		"        ",
		"  ####  ",
		" #      ",
		"  ####  ",
		"      # ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x0074, "\u0074", []string{ // "t"
		"        ",
		"  ##### ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"        "}},

	{0x0075, "\u0075", []string{ // "u"
		"        ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"  ####  ",
		"        "}},

	{0x0076, "\u0076", []string{ // "v"
		"        ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		"  #  #  ",
		"   ##   ",
		"        "}},

	{0x0077, "\u0077", []string{ // "w"
		"        ",
		" #    # ",
		" #    # ",
		" #    # ",
		" # ## # ",
		" ##  ## ",
		" #    # ",
		"        "}},

	{0x0078, "\u0078", []string{ // "x"
		"        ",
		" #    # ",
		"  #  #  ",
		"   ##   ",
		"   ##   ",
		"  #  #  ",
		" #    # ",
		"        "}},

	{0x0079, "\u0079", []string{ // "y"
		"        ",
		"  #   # ",
		"   # #  ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"        "}},

	{0x007A, "\u007A", []string{ // "z"
		"        ",
		" ###### ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" ###### ",
		"        "}},

	{0x007B, "\u007B", []string{ // "{"
		"  ###   ",
		" #      ",
		" #      ",
		"##      ",
		" #      ",
		" #      ",
		"  ###   ",
		"        "}},

	{0x007C, "\u007C", []string{ // "|"
		"   #    ",
		"   #    ",
		"   #    ",
		"        ",
		"   #    ",
		"   #    ",
		"   #    ",
		"        "}},

	{0x007D, "\u007D", []string{ // "}"
		"  ###   ",
		"     #  ",
		"     #  ",
		"     ## ",
		"     #  ",
		"     #  ",
		"  ###   ",
		"        "}},

	{0x007E, "\u007E", []string{ // "~"
		" ##     ",
		"#  #  # ",
		"    ##  ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        "}},

	{0x007F, "\u007F", []string{ // "[DEL]"
		"# # # # ",
		" # # #  ",
		"# # # # ",
		" # # #  ",
		"# # # # ",
		" # # #  ",
		"# # # # ",
		"        "}},

}
