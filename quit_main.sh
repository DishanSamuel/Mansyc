#!/bin/bash

./quitser.sh
./compare.sh

if [ ! -s "filter_user.txt" ]; then
  echo "All unnecessary services are already stopped."
  exit 0
else
  while read -r word; do
    word="${word%%.*}"

    if [ "$word" == "docker" ]; then
      sudo systemctl stop docker.socket
    fi

    sudo systemctl stop "$word"
    echo "$word stopped"
  done < filter_user.txt
fi
