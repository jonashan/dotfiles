#!/bin/bash

# Destination directory
DEST_DIR="/usr/local/bin"

# Excluded files (space-separated list)
EXCLUDE_FILES="install.sh README.md"

for file in *; do
    # Check if the file should be excluded
    if [[ " ${EXCLUDE_FILES[@]} " =~ " ${file} " ]]; then
        echo "Excluding file: ${file}"
    else
        # Create symlink if it's not a directory
        if [ ! -d "$file" ]; then
            ln -s "$(pwd)/${file}" "${DEST_DIR}/${file}"
            echo "Created symlink: $(pwd)/${file} -> ${DEST_DIR}/${file}"
        fi
    fi
done
