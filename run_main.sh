#!/bin/bash

echo "1. START SERVICES"
echo "2. END SERVICES"

read -p "Enter The Optinon: " option

if [ "$option" == "1" ]; then
    echo "Starting Services"
    ./start_main.sh
elif [ "$option" == "2" ]; then
    echo "Ending Services"
    ./quit_main.sh
    rm filter_user.txt | rm user_curser.txt
else
    echo "option not entered or invalid option"
    exit
fi