#!/bin/sh

message()
{
  TITLE="Cannot start ENERGY_HOME"
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

if [ -z "$ENERGY_HOME" ] || [ ! -d "$ENERGY_HOME" ]; then
  message "No CEF found. Please make sure \$ENERGY_HOME point to valid installation."
  exit 1
fi

LIBCEF="$ENERGY_HOME/libcef.so"

if [ ! -e "$LIBCEF" ]; then
  message "No CEF libcef.so found. Please make sure \$ENERGY_HOME point to valid installation."
  exit 1
fi

# fix: linux arm: Error loading libcef.so
# temp LD_PRELOAD, Load libcef.so correctly
export LD_PRELOAD="$LIBCEF"

# go run xxx.go
exec go run $1