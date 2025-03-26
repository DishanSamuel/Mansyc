#!/bin/bash

./quitser.sh

./compare.sh

if [ ! -s "filter_user.txt" ]; then
  echo "all unnessary services are already stopped"
  exit

else
    while read -r word; do

        word="${word%%.*}"
        sudo systemctl stop $word 
        echo "$word stopped"
        
    done < filter_user.txt 
fi