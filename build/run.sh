#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="gql-server"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
export $(grep -v '^#' build/.env | xargs) && time go run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' build/.env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: $app\n\n"