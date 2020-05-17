#!/bin/sh

echo -e "Shell script started\n"

echo -e "COMPILE: ${COMPILE:=NO}"
echo -e "RUN: ${RUN:=NO}"
RAW=/home/sandbox/raw
EXIT_CODE=0

if [ "$COMPILE" == "YES" ] ;
then
  echo -e "Compiling\n"
  if ! gcc $RAW/code_file.c -o code_binary 2> $RAW/compiler.txt; then
    echo "Compile failed"
    exit 254
  fi
  echo -e "Compilation Successful\n"
fi
if [ "$RUN" == "YES" ];
then
  echo -e "Running\n"
  echo -e "==================================================\n"
  ./code_binary 1> $RAW/output.txt
  EXIT_CODE="$?"
  echo -e "==================================================\n"
  echo -e "Run complete exiting...\n"
fi
echo "Sandbox exited with status code $EXIT_CODE"
exit "$EXIT_CODE"

# https://www.tldp.org/LDP/abs/html/exitcodes.html
