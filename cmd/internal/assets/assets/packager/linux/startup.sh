#!/bin/sh

message()
{
  TITLE="Cannot start {{.EXECUTE}}"
  if [ -n "$(command -v zenity)" ]; then
    zenity --error --title="$TITLE" --text="$1" --no-wrap
  elif [ -n "$(command -v kdialog)" ]; then
    kdialog --error "$1" --title "$TITLE"
  elif [ -n "$(command -v notify-send)" ]; then
    notify-send "ERROR: $TITLE" "$1"
  elif [ -n "$(command -v xmessage)" ]; then
    xmessage -center "ERROR: $TITLE: $1"
  else
    printf "ERROR: %s\n%s\n" "$TITLE" "$1"
  fi
}

# install path libcef.so
LIBCEF="{{.INSTALLPATH}}/libcef.so"

if [ ! -e "$LIBCEF" ]; then
  message "No CEF libcef.so found. Please make sure CEF point to valid installation."
  exit 1
fi

# fix: linux arm: Error loading libcef.so
# temp LD_PRELOAD, Load libcef.so correctly
export LD_PRELOAD="$LIBCEF"

# Execution File Name
STARTUP="{{.INSTALLPATH}}/{{.EXECUTE}}"

exec $STARTUP