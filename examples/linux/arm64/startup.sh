#!/bin/sh

# This is a solution to the CEF loading libcef.so error in the Linux ARM64 architecture
# The key lies in: export LD_PRELOAD="/you/cef/xxx/libcef.so"
# test: copy the shell to exe dir

# run: ./startup.sh demo

LIBCEF="$ENERGY_HOME/libcef.so"

echo $LIBCEF

# fix: linux arm: Error loading libcef.so
# temp LD_PRELOAD, Load libcef.so correctly
export LD_PRELOAD="$LIBCEF"

# Execution File Name
STARTUP="$PWD/$1"

exec $STARTU