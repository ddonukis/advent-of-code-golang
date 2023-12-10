#!/bin/bash

# Get the first argument passed to the script
folder_name=$1

# Create the subfolder relative to the script location
if [ -d "$(dirname "$0")/$folder_name" ]; then
    echo "$folder_name already exists"
    exit 1
fi
new_folder="$(dirname "$0")/$folder_name"
mkdir "$new_folder"


touch "$(dirname "$0")/$folder_name/example.txt"
touch "$(dirname "$0")/$folder_name/input.txt"

# Remove hyphens from the folder name
go_file_name=$(echo "$folder_name" | tr -d '-')
touch "$(dirname "$0")/$folder_name/$go_file_name.go"

echo "$new_folder"
