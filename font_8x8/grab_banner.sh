#! /bin/bash
# Grab sysvbanner 8x7 Font

WIDTH="8"
HEIGHT="8"

chr() {
  echo $1 | awk '{printf("%c", $1)}'
}

addspace() {
  echo -n "$1"
  LEN=`echo -n "$1" | wc -c`
  while [ "$LEN" -lt "$WIDTH" ]
  do
    echo -n " "
    LEN=$(($LEN + 1))
  done
}

# head
cat << EOF
var Font = []struct {
	cod int      // index (ASCII/KOI8-R code)
	chr string   // UTF-8 code
	val []string // [8]
}{
EOF

# body
for i in `seq 0x20 0xFF`
do
  CH=`chr $i`

  STR="$CH"

  [ "$CH" == "\""  ] && STR="\\$CH"
  [ "$CH" == "\\"  ] && STR="\\\\"
  [ "$i" -eq "127" ] && STR="[DEL]"

  printf "\t{0x%04X, \"\\\u%04X\", []string{ // \"%s\"" $i $i "$STR"
  
  SEP="\n\t\t"
  CNT="0"
  
  while IFS= read LINE
  do
    LINE=`addspace "$LINE"`
    printf "$SEP\"$LINE\""
    SEP=",\n\t\t"
    CNT=$(($CNT + 1))
  done < <(banner "$CH")

  while [ "$CNT" -lt "$HEIGHT" ]
  do
    LINE=`addspace " "`
    printf "$SEP\"$LINE\""
    CNT=$(($CNT + 1))
  done

  printf "}},\n\n"
done

# tail
echo "}"

