#!/bin/bash


echo "                       _________ __                 __            ________        .__  __   "
echo "  _________.__. ______/   _____//  |______ ________/  |_          \_____  \  __ __|__|/  |_ "
echo " /  ___<   |  |/  ___/\_____  \\   __\__  \\_  __ \   __\  ______  /  / \  \|  |  \  \   __\ "
echo " \___ \ \___  |\___ \ /        \|  |  / __ \|  | \/|  |   /_____/ /   \_/.  \  |  /  ||  |  "
echo "/____  >/ ____/____  >_______  /|__| (____  /__|   |__|           \_____\ \_/____/|__||__|  "
echo "     \/ \/         \/        \/           \/                             \__>               "

echo "
Welcome to my Bash script!
"


echo "1. START SERVICES"
echo "2. END SERVICES
"

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