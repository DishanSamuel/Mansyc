#!/bin/bash

systemctl list-units --type=service --state=running > curser.txt 
awk '{print $1}' curser.txt > new_curser.txt 
sed -i '1d' new_curser.txt 
head -n -6 new_curser.txt > pre_final_curser.txt
grep -v '^$' pre_final_curser.txt > user_curser.txt
cp user_curser.txt filter_user.txt
rm curser.txt | rm new_curser.txt | rm pre_final_curser.txt


