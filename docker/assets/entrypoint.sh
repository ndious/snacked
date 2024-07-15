#!/bin/bash

USER_PACKAGE=/assets/package.json
LOG=/var/log/assets

mkdir -p "$LOG"
echo "Log folder created"

cd /assets

if test -f "$USER_PACKAGE"; then
    echo "Installing dependencies"
    bun install
fi

SCRIPT=""

while getopts ":s:" option; do
   case $option in
      s) # Set bun script to execute
         SCRIPT=$OPTARG;;
   esac
done
echo "bun script to run $SCRIPT"
if test -n "$SCRIPT"; then
    cd /assets
    bun run "$SCRIPT" > "$LOG"/user-script.log 2> "$LOG"/user-script.error.log &
    echo "bun script $SCRIPT running"
fi

cd /usr/src/app
bun run statics
