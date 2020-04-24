#!/bin/sh
#
# A helper script to implement restart_container when the docker runtime isn't available.
#
# Usage:
#   Copy start.sh and restart.sh to your container working dir.
#
#   Make your container entrypoint:
#   ./start.sh path-to-binary [args]
#
#   To restart the container:
#   ./restart.sh

set -eu

process_id=""

trap quit TERM INT

quit() {
  if [ -n "group_id" ]; then
    # Kill the whole group
    kill -group_id
  fi
}

while true; do
    rm -f restart.txt

    setsid "$@" &
    group_id=$!
    echo "$group_id" > group_id.txt
    set +e
    wait $group_id
    EXIT_CODE=$?
    set -e
    if [ ! -f restart.txt ]; then
        echo "Exiting with code $EXIT_CODE"
        exit $EXIT_CODE
    fi
    echo "Restarting"
done
