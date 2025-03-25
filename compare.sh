#!/bin/bash

file1="normalser.txt" 
file2="filter_user.txt"  

while read -r word; do
    # Check if the word exists in file2
    if grep -qw "$word" "$file2"; then
        # Remove the word from file2
        sed -i "s/\b$word\b//g" "$file2"

    fi
    sed -i '/^$/d' "$file2"

done < "$file1"  # Ensure we read from file1
