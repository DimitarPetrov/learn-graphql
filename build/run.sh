#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="gql-server"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: %s\n" "$app"
# Set all ENV vars for the server to run
export $(grep -v '^#' build/.env | xargs) && time go run $src
printf "\nStopped running: %s\n\n" "$app"