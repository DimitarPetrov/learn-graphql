#!/bin/sh
app="gql-server"

printf "\nStart running: %s\n" "$app"
# Set all ENV vars for the server to run
export $(grep -v '^#' build/.env | xargs)
time /$GOPATH/bin/realize start run
# This should unset all the ENV vars, just in case.
printf "\nStopped running: %s\n\n" "$app"