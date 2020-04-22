#!/bin/bash

echo -e "Shell script started\n"

echo -e "COMPILE: ${COMPILE:=NO}"
echo -e "RUN: ${RUN:=NO}"
EXIT_CODE=0

if [ "$COMPILE" == "YES" ] ;
then
  echo -e "Compiling\n"
  if ! gcc solution.c -o solution_binary 2> compiler.txt; then
    echo "Compile failed"
    exit 254
  fi
  echo -e "Compiled\n"
fi
if [ "$RUN" == "YES" ];
then
  echo -e "Running\n"
  echo -e "==================================================\n"
  ./solution_binary 1> output.txt
  EXIT_CODE="$?"
  echo -e "==================================================\n"
  echo -e "Run complete exiting...\n"
fi
echo "Sandbox exited with status code $EXIT_CODE"
exit "$EXIT_CODE"

# https://www.tldp.org/LDP/abs/html/exitcodes.html