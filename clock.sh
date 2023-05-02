#! /bin/bash

#DOT="\u2588"
#DOT="\u25A0"
DOT="O"

DOT=`echo -e "$DOT"`

while true
do
  clear
  #DATE=`date '+%m-%d'`
  TIME=`date '+%H:%M:%S'`
  #subanner -d "$DOT" "$DATE" "$TIME"
  subanner -d "$DOT" "$TIME"
  sleep 1
done

