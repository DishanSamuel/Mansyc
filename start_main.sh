#!/bin/bash

if [ ! -s "startup.txt" ]; then
  echo "no startup files entered"
  exit

else
    while read -r word; do

        word="${word%%.*}"
        sudo systemctl start $word 
        echo "$word started"
        
    done < startup.txt 
fi